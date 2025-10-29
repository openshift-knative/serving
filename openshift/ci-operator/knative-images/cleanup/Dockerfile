# DO NOT EDIT! Generated Dockerfile for pkg/cleanup/cmd/cleanup.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.24-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./pkg/cleanup/cmd/cleanup

FROM $GO_RUNTIME

ARG VERSION=knative-v1.18

COPY --from=builder /usr/bin/main /ko-app/cleanup
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-cleanup-rhel9-container" \
      name="openshift-serverless-1/serving-cleanup-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Cleanup" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Cleanup" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Cleanup" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Cleanup" \
      io.openshift.tags="cleanup" \
      vendor="Red Hat, Inc." \
      release=$VERSION \
      cpe="cpe:/a:redhat:openshift_serverless:1.38::el9"

ENTRYPOINT ["/ko-app/cleanup"]
