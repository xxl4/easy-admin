export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w
PROJECT:=easy-admin
VERSION := 1.1.0


.PHONY: build

all: 
	make build-ui
	make build

build-ui:
	@echo "build node start"
	cd ./ui/ && npm run build:prod

build:
	CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -a -installsuffix "" -o easy-admin .

# make build-linux
build-linux:
	make build-ui
	make build
	@docker build -t easy-admin:$(VERSION) .
	@echo "build successful"

build-sqlite:
	go build -tags sqlite3 -ldflags="$(LDFLAGS)" -a -installsuffix -o easy-admin .

test:
	go test
