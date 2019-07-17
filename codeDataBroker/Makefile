SHELL=bash
# Expands to list this project's go packages, excluding the vendor folder
PACKAGES = $$(go list ./... | grep -v /vendor/)

ifeq ($(shell uname -s),Darwin)
endif

build: builddir
	go build -a -o build/producer ./cmd/producer
	go build -a -o build/broker ./cmd/broker

producer:
	rm -rf build/producer
	go build -a -o build/producer ./cmd/producer

broker:
	rm -rf build/broker
	go build -a -o build/broker ./cmd/broker

consumer:
	rm -rf build/consumer
	go build -a -o build/consumer ./cmd/consumer

#MY_IP Variable used in docker-compose is mentioned in .env
#Ip can be detched through command `ifconfig -a| grep -A1 enp0| grep inet | cut -d ' ' -f10`. TODO Automate this. Can be used like MY_IP=`<comamnd>` docker-compose up 
kafkaCluster:
	docker-compose up -d 

kafkaClusterDown:
	docker-compose down

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



