all: protoc build

protoc:
		protoc -I$$GOPATH/src --go_out=plugins=micro:$$GOPATH/src $$PWD/proto/*.proto
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  .
