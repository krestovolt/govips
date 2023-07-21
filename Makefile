all: deps build test

deps: FORCE
	CGO_CFLAGS_ALLOW=-Xpreprocessor go get ./...

build: FORCE
	CGO_CFLAGS_ALLOW=-Xpreprocessor go build ./vips

test: FORCE
	SKIP_TEST_INIT=1 CGO_CFLAGS_ALLOW=-Xpreprocessor go test -v ./...
	CGO_CFLAGS_ALLOW=-Xpreprocessor go test -v -run '^TestInitConfig$$' ./...

FORCE:
