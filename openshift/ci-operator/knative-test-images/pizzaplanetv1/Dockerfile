# Do not edit! This file was generated via Makefile
FROM registry.ci.openshift.org/openshift/release:golang-1.21 as builder

COPY . .
RUN make install test-install

FROM registry.access.redhat.com/ubi8/ubi-minimal
USER 65532

COPY --from=builder /go/bin/pizzaplanetv1 /ko-app/pizzaplanetv1
ENTRYPOINT ["/ko-app/pizzaplanetv1"]
