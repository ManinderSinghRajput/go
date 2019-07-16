SHELL=bash
# Expands to list this project's go packages, excluding the vendor folder
PACKAGES = $$(go list ./... | grep -v /vendor/)

ifeq ($(shell uname -s),Darwin)
endif

build: builddir
	go build -a -o build/producer ./cmd/producer
	go build -a -o build/broker ./cmd/broker

producer:
	go build -a -o build/producer ./cmd/producer

broker:
	go build -a -o build/broker ./cmd/broker

consumer:
	go build -a -o build/consumer ./cmd/consumer

builddir:
	if [ ! -d build ]; then mkdir build; fi

run: build up

vet:
	go vet $(PACKAGES)

lint:
	golint -set_exit_status $(PACKAGES)

clean:
	rm -rf build/*

fmt:
	go fmt $(PACKAGES)

# Run this before committing.
prep: fmt vet clean build

vendor.update:
	dep ensure --update

tools: buildtools
	go get -u github.com/golang/dep/cmd/dep

buildtools:
	go get -u github.com/tebeka/go2xunit
	go get -u golang.org/x/tools/cmd/cover
	go get -u golang.org/x/lint/golint



