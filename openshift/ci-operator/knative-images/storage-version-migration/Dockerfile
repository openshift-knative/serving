# DO NOT EDIT! Generated Dockerfile for cmd/storage-version-migration.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.25-openshift-4.21
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/storage-version-migration

FROM $GO_RUNTIME

ARG VERSION=knative-nightly

COPY --from=builder /usr/bin/main /ko-app/storage-version-migration
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-serving-storage-version-migration-rhel9-container" \
      name="openshift-serverless-1/serving-storage-version-migration-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Serving Storage Version Migration" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Serving Storage Version Migration" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Serving Storage Version Migration" \
      io.k8s.description="Red Hat OpenShift Serverless Serving Storage Version Migration" \
      io.openshift.tags="storage-version-migration" \
      vendor="Red Hat, Inc." \
      release=$VERSION \
      cpe="cpe:/a:redhat:openshift_serverless:::el9"

ENTRYPOINT ["/ko-app/storage-version-migration"]
