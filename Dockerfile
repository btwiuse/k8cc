FROM golang:latest as builder

RUN CGO_ENABLED=0 GOOS=linux go get github.com/btwiuse/k8cc
# CMD ["/go/bin/k8cc", "api", "-logtostderr=true"]
# CMD ["/go/bin/k8cc", "controller", "-logtostderr=true"]

# FROM scratch
# COPY --from=builder /go/bin/k8cc /k8cc
# CMD ["/k8cc", "api", "-logtostderr=true"]
