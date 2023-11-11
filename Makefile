export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
LDFLAGS := -s -w
# application name
PROJECT:=easy-admin
# application version
VERSION := 1.2.0
# application url
URL := https://github.com/nicelizhi/easy-admin

.PHONY: build

all: 
	make build-ui
	make build

# build vue ui
build-ui:
	@echo "build node start"
	cd ./ui/ && npm run build:prod:base

# build go application
build:
	CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -a -installsuffix "" -o $(PROJECT) .

# make build-linux
build-linux:
	make build-ui
	make build
	@docker build -t $(PROJECT):$(VERSION) .
	@echo "build successful"

# build sql go application version
build-sqlite:
	go build -tags sqlite3 -ldflags="$(LDFLAGS)" -a -installsuffix -o $(PROJECT) .

clean:
	$(GOCLEAN)
	rm ./$(PROJECT)

test:
	$(GOTEST)

restart:
	make stop
	make start

.PHONY: start
start:
	nohup ./$(PROJECT) server -c=config/settings.dev_steve.yml >> acc.txt &
	ps aux | grep "$(PROJECT)"

stop:
	pkill $(PROJECT)

# debug file
debug:
	dlv debug $(file)