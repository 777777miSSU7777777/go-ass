.PHONY: fmt
fmt:
	goimports -w ./

.PHONY: build
build:
	go build