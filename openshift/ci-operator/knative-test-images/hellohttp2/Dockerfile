# DO NOT EDIT! Generated Dockerfile for test/test_images/hellohttp2.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./test/test_images/hellohttp2

FROM $GO_RUNTIME

ARG VERSION=knative-v1.15

RUN microdnf install tzdata 

USER 65532

COPY --from=builder /usr/bin/main /ko-app/hellohttp2

LABEL \
      com.redhat.component="openshift-serverless-1-serving-test-test-images-hellohttp2-rhel8-container" \
      name="openshift-serverless-1/serving-test-test-images-hellohttp2-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Test Test Images Hellohttp2" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Test Test Images Hellohttp2" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Test Test Images Hellohttp2"

ENTRYPOINT ["/ko-app/hellohttp2"]
