# DO NOT EDIT! Generated Dockerfile for test/test_images/hellohttp2.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.23-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./test/test_images/hellohttp2

FROM $GO_RUNTIME

ARG VERSION=knative-v1.17

COPY --from=builder /usr/bin/main /ko-app/hellohttp2
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-test-test-images-hellohttp2-rhel9-container" \
      name="openshift-serverless-1/serving-test-test-images-hellohttp2-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Test Test Images Hellohttp2" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Test Test Images Hellohttp2" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Test Test Images Hellohttp2" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Test Test Images Hellohttp2" \
      io.openshift.tags="test-test-images-hellohttp2" \
      vendor="Red Hat, Inc." \
      release=$VERSION \
      cpe="cpe:/a:redhat:openshift_serverless:1.37::el9"

ENTRYPOINT ["/ko-app/hellohttp2"]
