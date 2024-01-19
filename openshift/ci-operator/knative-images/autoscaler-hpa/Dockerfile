# Do not edit! This file was generated via Makefile
FROM registry.ci.openshift.org/openshift/release:golang-1.19 as builder

COPY . .
RUN make install test-install

FROM registry.access.redhat.com/ubi8/ubi-minimal
USER 65532

COPY ./cmd/autoscaler-hpa/kodata /var/run/ko
COPY --from=builder /go/bin/autoscaler-hpa /ko-app/autoscaler-hpa
ENTRYPOINT ["/ko-app/autoscaler-hpa"]
