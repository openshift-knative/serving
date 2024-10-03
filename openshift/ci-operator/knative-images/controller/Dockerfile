# DO NOT EDIT! Generated Dockerfile for cmd/controller.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/controller

FROM $GO_RUNTIME

ARG VERSION=knative-nightly

COPY --from=builder /usr/bin/main /ko-app/controller

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-controller-rhel8-container" \
      name="openshift-serverless-1/serving-controller-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Controller" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Controller" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Controller" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Controller" \
      io.openshift.tags="controller"

ENTRYPOINT ["/ko-app/controller"]
