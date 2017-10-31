install:
	go get -u github.com/golang/dep/cmd/dep
.PHONY: install

init:
	dep ensure
.PHONY: init

test:
	go test -v (go list ./... | grep -v vendor)
.PHONY: test
