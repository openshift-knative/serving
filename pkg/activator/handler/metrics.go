/*
Copyright 2020 The Knative Authors

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

package handler

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var scopeName = "knative.dev/serving/pkg/activator"

type ccMetrics struct {
	requestCC     metric.Float64Gauge
	usePrometheus bool
}

func newMetrics(mp metric.MeterProvider, usePrometheus bool) *ccMetrics {
	m := ccMetrics{usePrometheus: usePrometheus}

	if usePrometheus {
		registerPromMetrics()
		return &m
	}

	provider := mp
	if provider == nil {
		provider = otel.GetMeterProvider()
	}

	meter := provider.Meter(scopeName)

	var err error
	m.requestCC, err = meter.Float64Gauge(
		"kn.revision.request.concurrency",
		metric.WithDescription("Concurrent requests that are routed to Activator"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		panic(err)
	}

	return &m
}

var (
	promRegisterOnce sync.Once

	durationBuckets = []float64{0.005, 0.01, 0.025, 0.05, 0.075, 0.1, 0.25, 0.5, 0.75, 1, 2.5, 5, 7.5, 10}

	// OTel SDK default histogram boundaries (reader.go DefaultAggregationSelector for InstrumentKindHistogram).
	bodySizeBuckets = []float64{0, 5, 10, 25, 50, 75, 100, 250, 500, 750, 1000, 2500, 5000, 7500, 10000}

	// Labels match the OTel-to-Prometheus naming convention (dots → underscores)
	// used by the v1.21 otelhttp metrics.
	httpMetricLabels = []string{
		"http_request_method",
		"http_response_status_code",
		"network_protocol_name",
		"network_protocol_version",
		"url_scheme",
		"k8s_namespace_name",
		"kn_service_name",
		"kn_configuration_name",
		"kn_revision_name",
	}

	concurrencyLabels = []string{
		"k8s_namespace_name",
		"kn_service_name",
		"kn_configuration_name",
		"kn_revision_name",
	}

	serverRequestDuration  *prometheus.HistogramVec
	serverRequestBodySize  *prometheus.HistogramVec
	serverResponseBodySize *prometheus.HistogramVec
	clientRequestDuration  *prometheus.HistogramVec
	clientRequestBodySize  *prometheus.HistogramVec
	requestConcurrency     *prometheus.GaugeVec
)

func registerPromMetrics() {
	promRegisterOnce.Do(func() {
		serverRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_server_request_duration_seconds",
			Help:    "Duration of HTTP server requests.",
			Buckets: durationBuckets,
		}, httpMetricLabels)

		serverRequestBodySize = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_server_request_body_size_bytes",
			Help:    "Size of HTTP server request bodies.",
			Buckets: bodySizeBuckets,
		}, httpMetricLabels)

		serverResponseBodySize = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_server_response_body_size_bytes",
			Help:    "Size of HTTP server response bodies.",
			Buckets: bodySizeBuckets,
		}, httpMetricLabels)

		clientRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_client_request_duration_seconds",
			Help:    "Duration of HTTP client requests.",
			Buckets: durationBuckets,
		}, httpMetricLabels)

		clientRequestBodySize = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_client_request_body_size_bytes",
			Help:    "Size of HTTP client request bodies.",
			Buckets: bodySizeBuckets,
		}, httpMetricLabels)

		requestConcurrency = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "kn_revision_request_concurrency",
			Help: "Concurrent requests that are routed to Activator.",
		}, concurrencyLabels)

		prometheus.MustRegister(
			serverRequestDuration,
			serverRequestBodySize,
			serverResponseBodySize,
			clientRequestDuration,
			clientRequestBodySize,
			requestConcurrency,
		)
	})
}

// DeleteRevisionMetrics removes all metrics associated with the given revision.
func DeleteRevisionMetrics(namespace, serviceName, configName, revisionName string) {
	labels := prometheus.Labels{
		"k8s_namespace_name":    namespace,
		"kn_service_name":       serviceName,
		"kn_configuration_name": configName,
		"kn_revision_name":      revisionName,
	}
	serverRequestDuration.DeletePartialMatch(labels)
	serverRequestBodySize.DeletePartialMatch(labels)
	serverResponseBodySize.DeletePartialMatch(labels)
	clientRequestDuration.DeletePartialMatch(labels)
	clientRequestBodySize.DeletePartialMatch(labels)
	requestConcurrency.DeletePartialMatch(labels)
}
