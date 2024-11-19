# DO NOT EDIT! Generated Dockerfile for cmd/autoscaler.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/autoscaler

FROM $GO_RUNTIME

ARG VERSION=knative-v1.15

COPY --from=builder /usr/bin/main /ko-app/autoscaler
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-autoscaler-rhel8-container" \
      name="openshift-serverless-1/serving-autoscaler-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Autoscaler" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Autoscaler" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Autoscaler" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Autoscaler" \
      io.openshift.tags="autoscaler"

ENTRYPOINT ["/ko-app/autoscaler"]
