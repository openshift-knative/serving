/*
Copyright 2025 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sharedmain

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/prometheus/otlptranslator"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	otelmetric "go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"knative.dev/pkg/changeset"
	"knative.dev/pkg/observability/metrics"
	promserver "knative.dev/pkg/observability/metrics/prometheus"
	"knative.dev/pkg/observability/semconv"
	"knative.dev/pkg/observability/tracing"
	"knative.dev/pkg/system"
	servingmetrics "knative.dev/serving/pkg/metrics"
	"knative.dev/serving/pkg/networking"
)

// observabilityMeterProvider combines the OTel metric provider interface with a Shutdown method.
type observabilityMeterProvider interface {
	otelmetric.MeterProvider
	Shutdown(context.Context) error
}

func SetupObservabilityOrDie(
	ctx context.Context,
	cfg *config,
	logger *zap.SugaredLogger,
) (observabilityMeterProvider, *tracing.TracerProvider) {
	r := res(logger, cfg)

	var mp observabilityMeterProvider
	if cfg.Observability.RequestMetrics.Protocol == metrics.ProtocolPrometheus {
		mp = buildPrometheusProvider(cfg, r, logger)
	} else {
		meterProvider, err := metrics.NewMeterProvider(
			ctx,
			cfg.Observability.RequestMetrics,
			sdkmetric.WithResource(r),
		)
		if err != nil {
			logger.Fatalw("Failed to setup meter provider", zap.Error(err))
		}
		mp = meterProvider
	}

	otel.SetMeterProvider(mp)

	err := runtime.Start(
		runtime.WithMinimumReadMemStatsInterval(cfg.Observability.Runtime.ExportInterval),
		runtime.WithMeterProvider(mp),
	)
	if err != nil {
		logger.Fatalw("Failed to start runtime metrics", zap.Error(err))
	}

	tracerProvider, err := tracing.NewTracerProvider(
		ctx,
		cfg.Observability.Tracing,
		trace.WithResource(r),
	)
	if err != nil {
		logger.Fatalw("Failed to setup trace provider", zap.Error(err))
	}

	otel.SetTextMapPropagator(tracing.DefaultTextMapPropagator())
	otel.SetTracerProvider(tracerProvider)

	return mp, tracerProvider
}

// buildPrometheusProvider creates a Prometheus-backed meter provider that exposes
// the serving-specific resource attributes (service, configuration, revision) as
// constant Prometheus labels on every metric.
func buildPrometheusProvider(cfg *config, r *resource.Resource, logger *zap.SugaredLogger) observabilityMeterProvider {
	endpoint := cfg.Observability.RequestMetrics.Endpoint
	if endpoint == "" {
		endpoint = fmt.Sprintf(":%d", networking.UserQueueMetricsPort)
	}

	reader, err := otelprom.New(
		otelprom.WithTranslationStrategy(otlptranslator.UnderscoreEscapingWithSuffixes),
		otelprom.WithResourceAsConstantLabels(attribute.NewAllowKeysFilter(
			attribute.Key(servingmetrics.ServiceNameKey),
			attribute.Key(servingmetrics.ConfigurationNameKey),
			attribute.Key(servingmetrics.RevisionNameKey),
		)),
	)
	if err != nil {
		logger.Fatalw("Failed to create prometheus exporter", zap.Error(err))
	}

	host, port, err := net.SplitHostPort(endpoint)
	if err != nil {
		logger.Fatalw("Invalid prometheus endpoint", zap.String("endpoint", endpoint), zap.Error(err))
	}

	server, err := promserver.NewServer(
		promserver.WithHost(host),
		promserver.WithPort(port),
	)
	if err != nil {
		logger.Fatalw("Failed to create prometheus server", zap.Error(err))
	}
	go server.ListenAndServe()

	sdkProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(reader),
		sdkmetric.WithResource(r),
	)

	return &prometheusProvider{
		MeterProvider: sdkProvider,
		shutdownFns:   []func(context.Context) error{sdkProvider.Shutdown, server.Shutdown},
	}
}

type prometheusProvider struct {
	otelmetric.MeterProvider
	shutdownFns []func(context.Context) error
}

func (p *prometheusProvider) Shutdown(ctx context.Context) error {
	errs := make([]error, 0, len(p.shutdownFns))
	for _, fn := range p.shutdownFns {
		errs = append(errs, fn(ctx))
	}
	return errors.Join(errs...)
}

func res(logger *zap.SugaredLogger, cfg *config) *resource.Resource {
	podName := system.PodName()

	serviceName := cmp.Or(
		os.Getenv("OTEL_SERVICE_NAME"),
		os.Getenv("SERVING_SERVICE"),
		os.Getenv("SERVING_CONFIGURATION"),
		os.Getenv("SERVING_REVISION"),

		// I always expect SERVING_REVISION to be set but in case it's
		// not fallback on pod name
		podName,
	)

	attrs := []attribute.KeyValue{
		semconv.K8SContainerName("queue-proxy"),
		semconv.K8SNamespaceName(cfg.ServingNamespace),
		semconv.K8SPodName(podName),
		semconv.ServiceVersion(changeset.Get()),
		semconv.ServiceName(serviceName),
		semconv.ServiceInstanceID(podName),
	}

	if val := os.Getenv("SERVING_SERVICE"); val != "" {
		attrs = append(attrs, servingmetrics.ServiceNameKey.With(val))
	}
	if val := os.Getenv("SERVING_CONFIGURATION"); val != "" {
		attrs = append(attrs, servingmetrics.ConfigurationNameKey.With(val))
	}
	if val := os.Getenv("SERVING_REVISION"); val != "" {
		attrs = append(attrs, servingmetrics.RevisionNameKey.With(val))
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			attrs...,
		),
	)
	if err != nil {
		logger.Fatalw("error merging otel resources", zap.Error(err))
	}

	return r
}
