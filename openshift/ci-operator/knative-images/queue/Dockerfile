# Do not edit! This file was generated via Makefile
FROM registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.21-openshift-4.16 as builder

COPY . .
RUN make install test-install

FROM registry.access.redhat.com/ubi8/ubi-minimal
USER 65532

COPY ./cmd/queue/kodata /var/run/ko
COPY --from=builder /go/bin/queue /ko-app/queue
ENTRYPOINT ["/ko-app/queue"]
