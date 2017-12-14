BINARY:=bora
VERSION=0.0.0
MONOVA:=$(shell which monova dot 2> /dev/null)

version:
ifdef MONOVA
override VERSION="$(shell monova)"
else
	$(info "Install monova (https://github.com/jsnjack/monova) to calculate version")
endif

build: version
	go build -ldflags="-X main.version=${VERSION}"

run: build
	./${BINARY}
