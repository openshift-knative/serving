# DO NOT EDIT! Generated Dockerfile for cmd/autoscaler.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.23-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/autoscaler

FROM $GO_RUNTIME

ARG VERSION=knative-v1.17

COPY --from=builder /usr/bin/main /ko-app/autoscaler
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-serving-autoscaler-rhel9-container" \
      name="openshift-serverless-1/kn-serving-autoscaler-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Autoscaler" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Autoscaler" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Autoscaler" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Autoscaler" \
      io.openshift.tags="autoscaler" \
      vendor="Red Hat, Inc." \
      release=$VERSION \
      cpe="cpe:/a:redhat:openshift_serverless:1.37::el9"

ENTRYPOINT ["/ko-app/autoscaler"]
