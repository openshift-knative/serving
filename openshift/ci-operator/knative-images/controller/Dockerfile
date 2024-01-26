# Do not edit! This file was generated via Makefile
FROM registry.ci.openshift.org/openshift/release:golang-1.21 as builder

COPY . .
RUN make install test-install

FROM registry.access.redhat.com/ubi8/ubi-minimal
USER 65532

COPY ./cmd/controller/kodata /var/run/ko
COPY --from=builder /go/bin/controller /ko-app/controller
ENTRYPOINT ["/ko-app/controller"]
