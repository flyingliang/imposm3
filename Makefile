.PHONY: test all build clean

PROTOFILES=$(shell find . -name \*.proto)
PBGOFILES=$(patsubst %.proto,%.pb.go,$(PROTOFILES))
GOFILES=$(shell find . \( -name \*.go ! -name version.go \) )

# for protoc-gen-go
export PATH := $(GOPATH)/bin:$(PATH)

GOLDFLAGS=-ldflags '-r $${ORIGIN}/lib'


BUILD_DATE=$(shell date +%Y%m%d)
BUILD_REV=$(shell git rev-parse --short HEAD)
BUILD_VERSION=dev-$(BUILD_DATE)-$(BUILD_REV)

all: build test

imposm3: $(GOFILES) $(PROTOFILES)
	@sed -i='' 's/buildVersion = ".*"/buildVersion = "$(BUILD_VERSION)"/' version.go
	go build $(GOLDFLAGS)
	@sed -i='' 's/buildVersion = ".*"/buildVersion = ""/' version.go

build: imposm3

clean:
	rm -f imposm3
	(cd test && make clean)

test:
	(cd test && make test)

%.pb.go: %.proto
	protoc --go_out=. $^
