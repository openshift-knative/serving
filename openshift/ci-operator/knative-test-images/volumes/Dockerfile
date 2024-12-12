# DO NOT EDIT! Generated Dockerfile for test/test_images/volumes.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./test/test_images/volumes

FROM $GO_RUNTIME

ARG VERSION=knative-nightly

COPY --from=builder /usr/bin/main /ko-app/volumes
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-test-test-images-volumes-rhel8-container" \
      name="openshift-serverless-1/serving-test-test-images-volumes-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Test Test Images Volumes" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Test Test Images Volumes" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Test Test Images Volumes" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Test Test Images Volumes" \
      io.openshift.tags="test-test-images-volumes"

ENTRYPOINT ["/ko-app/volumes"]
