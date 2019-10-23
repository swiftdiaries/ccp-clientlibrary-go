GOPATH ?= $(HOME)/go
TAG ?= 0.1
OPSYS ?= darwin

build: go-install build-ccpctl

go-install:
	go mod vendor
	go mod download
build-ccpctl:
	env GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ${GOPATH}/src/github.com/swiftdiaries/ccp-clientlibrary-go/bin/ccpctl-linux ${GOPATH}/src/github.com/swiftdiaries/ccp-clientlibrary-go/cmd/ccpctl/ccpctl.go
	env GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o ${GOPATH}/src/github.com/swiftdiaries/ccp-clientlibrary-go/bin/ccpctl-darwin ${GOPATH}/src/github.com/swiftdiaries/ccp-clientlibrary-go/cmd/ccpctl/ccpctl.go	
test: 
clean:
tar:
release: tar
