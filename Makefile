export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w
PROJECT:=easy-admin


.PHONY: build

all: 
	make build-ui
	make build

build-ui:
	@echo "build node start"
	cd ./admin-ui/ && npm run build:prod

build:
	CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -a -installsuffix "" -o easy-admin .

# make build-linux
build-linux:
	@docker build -t go-admin:latest .
	@echo "build successful"

build-sqlite:
	go build -tags sqlite3 -ldflags="$(LDFLAGS)" -a -installsuffix -o easy-admin .
