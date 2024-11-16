# DO NOT EDIT! Generated Dockerfile for test/test_images/slowstart.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /go/src/test/test_images/slowstart
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./test/test_images/slowstart

FROM $GO_RUNTIME

ARG VERSION=knative-nightly

COPY --from=builder /usr/bin/main /ko-app/slowstart

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-test-test-images-slowstart-rhel8-container" \
      name="openshift-serverless-1/serving-test-test-images-slowstart-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Test Test Images Slowstart" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Test Test Images Slowstart" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Test Test Images Slowstart" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Test Test Images Slowstart" \
      io.openshift.tags="test-test-images-slowstart"

ENTRYPOINT ["/ko-app/slowstart"]
