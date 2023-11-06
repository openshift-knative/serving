# Do not edit! This file was generated via Makefile
FROM registry.ci.openshift.org/openshift/release:golang-1.19 as builder

COPY . .

RUN ./openshift/in-docker-patch.sh
RUN make perf-install

FROM openshift/origin-base
USER 65532

COPY --from=builder /go/bin/real-traffic-test /ko-app/real-traffic-test
ENTRYPOINT ["/ko-app/real-traffic-test"]
