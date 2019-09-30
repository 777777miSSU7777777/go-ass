.PHONY: fmt
fmt:
	goimports -w ./

.PHONY: build
build:
	go build

.PHONY: run
run: 
	docker-compose up --build