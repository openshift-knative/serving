# DO NOT EDIT! Generated Dockerfile for cmd/default-domain.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.24-nofips-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/default-domain

FROM $GO_RUNTIME

ARG VERSION=knative-nightly

COPY --from=builder /usr/bin/main /ko-app/default-domain
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-default-domain-rhel8-container" \
      name="openshift-serverless-1/serving-default-domain-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Default Domain" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Default Domain" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Default Domain" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Default Domain" \
      io.openshift.tags="default-domain"

ENTRYPOINT ["/ko-app/default-domain"]
