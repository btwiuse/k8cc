all:
	go build ./cmd/k8cc-controller
	go build ./cmd/k8cc-api

test:
	go test ./...
	./hack/verify-codegen.sh
	gometalinter --vendor --skip=pkg/client --skip=pkg/apis --deadline=5m ./...

gen:
	./hack/update-codegen.sh
	go generate ./...

docker:
	docker build -t "btwiuse/k8cc:latest" -f Dockerfile .

.PHONY: all test gen docker
