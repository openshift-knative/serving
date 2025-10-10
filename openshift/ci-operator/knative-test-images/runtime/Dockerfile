# DO NOT EDIT! Generated Dockerfile for test/test_images/runtime.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.23-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./test/test_images/runtime

FROM $GO_RUNTIME

ARG VERSION=knative-v1.17

COPY --from=builder /usr/bin/main /ko-app/runtime
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-test-test-images-runtime-rhel9-container" \
      name="openshift-serverless-1/serving-test-test-images-runtime-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Test Test Images Runtime" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Test Test Images Runtime" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Test Test Images Runtime" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Test Test Images Runtime" \
      io.openshift.tags="test-test-images-runtime" \
      cpe="cpe:/a:redhat:openshift_serverless:1.37::el9"

ENTRYPOINT ["/ko-app/runtime"]
