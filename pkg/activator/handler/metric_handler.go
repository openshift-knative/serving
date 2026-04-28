/*
Copyright 2019 The Knative Authors

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
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"knative.dev/serving/pkg/apis/serving"
	pkghttp "knative.dev/serving/pkg/http"
	"knative.dev/serving/pkg/metrics"
)

// NewMetricHandler creates a handler that either adds serving attributes
// to the otelhttp labeler (OTel path) or records server-side HTTP metrics
// directly to Prometheus histograms (Prometheus path).
func NewMetricHandler(podName string, next http.Handler, usePrometheus bool) *MetricHandler {
	if usePrometheus {
		registerPromHTTPMetrics()
	}
	return &MetricHandler{
		nextHandler:   next,
		podName:       podName,
		usePrometheus: usePrometheus,
	}
}

// MetricHandler is a handler that records request metrics.
type MetricHandler struct {
	podName       string
	nextHandler   http.Handler
	usePrometheus bool
}

func (h *MetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.usePrometheus {
		h.serveHTTPPrometheus(w, r)
	} else {
		h.serveHTTPOTel(w, r)
	}
}

func (h *MetricHandler) serveHTTPOTel(w http.ResponseWriter, r *http.Request) {
	rev := RevisionFrom(r.Context())

	serviceName := rev.Labels[serving.ServiceLabelKey]
	configurationName := rev.Labels[serving.ConfigurationLabelKey]

	labeler, _ := otelhttp.LabelerFromContext(r.Context())
	labeler.Add(
		metrics.ServiceNameKey.With(serviceName),
		metrics.ConfigurationNameKey.With(configurationName),
		metrics.RevisionNameKey.With(rev.Name),
		metrics.K8sNamespaceKey.With(rev.Namespace),
	)

	h.nextHandler.ServeHTTP(w, r)
}

func (h *MetricHandler) serveHTTPPrometheus(w http.ResponseWriter, r *http.Request) {
	rev := RevisionFrom(r.Context())

	svc := rev.Labels[serving.ServiceLabelKey]
	cfg := rev.Labels[serving.ConfigurationLabelKey]
	ns := rev.Namespace
	revName := rev.Name

	method := standardizeMethod(r.Method)
	protoName, protoVersion := protoInfo(r.Proto)
	scheme := urlScheme(r.TLS != nil)

	rr := pkghttp.NewResponseRecorder(w, http.StatusOK)
	start := time.Now()

	defer func() {
		duration := time.Since(start).Seconds()
		statusCode := strconv.Itoa(rr.ResponseCode)

		// Label order must match httpMetricLabels in metrics.go:
		// http_request_method, http_response_status_code,
		// network_protocol_name, network_protocol_version, url_scheme,
		// k8s_namespace_name, kn_service_name, kn_configuration_name, kn_revision_name
		labels := []string{method, statusCode, protoName, protoVersion, scheme, ns, svc, cfg, revName}

		serverRequestDuration.WithLabelValues(labels...).Observe(duration)
		if r.ContentLength >= 0 {
			serverRequestBodySize.WithLabelValues(labels...).Observe(float64(r.ContentLength))
		}
		serverResponseBodySize.WithLabelValues(labels...).Observe(float64(rr.ResponseSize))
	}()

	h.nextHandler.ServeHTTP(rr, r)
}

func standardizeMethod(method string) string {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut,
		http.MethodPatch, http.MethodDelete, http.MethodConnect,
		http.MethodOptions, http.MethodTrace:
		return method
	default:
		return "_OTHER"
	}
}

func protoInfo(proto string) (name, version string) {
	switch proto {
	case "HTTP/1.0":
		return "http", "1.0"
	case "HTTP/1.1":
		return "http", "1.1"
	case "HTTP/2.0", "HTTP/2":
		return "http", "2"
	default:
		return "", ""
	}
}

func urlScheme(tls bool) string {
	if tls {
		return "https"
	}
	return "http"
}
