# DO NOT EDIT! Generated Dockerfile for cmd/autoscaler-hpa.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.24-nofips-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/autoscaler-hpa

FROM $GO_RUNTIME

ARG VERSION=knative-nightly

COPY --from=builder /usr/bin/main /ko-app/autoscaler-hpa
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-autoscaler-hpa-rhel8-container" \
      name="openshift-serverless-1/serving-autoscaler-hpa-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Autoscaler Hpa" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Autoscaler Hpa" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Autoscaler Hpa" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Autoscaler Hpa" \
      io.openshift.tags="autoscaler-hpa"

ENTRYPOINT ["/ko-app/autoscaler-hpa"]
