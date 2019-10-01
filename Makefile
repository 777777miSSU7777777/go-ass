.PHONY: fmt
fmt:
	goimports -w ./

.PHONY: build
build:
	go build

.PHONY: run
run: 
	[ -d $(HOME)/.goass/storage ] || mkdir -p $(HOME)/.goass/storage
	docker-compose up --build