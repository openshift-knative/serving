# DO NOT EDIT! Generated Dockerfile for cmd/schema-tweak.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/schema-tweak

FROM $GO_RUNTIME

ARG VERSION=knative-v1.17

COPY --from=builder /usr/bin/main /ko-app/schema-tweak
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-schema-tweak-rhel8-container" \
      name="openshift-serverless-1/serving-schema-tweak-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Schema Tweak" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Schema Tweak" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Schema Tweak" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Schema Tweak" \
      io.openshift.tags="schema-tweak"

ENTRYPOINT ["/ko-app/schema-tweak"]
