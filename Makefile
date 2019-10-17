.PHONY: default
default: all

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build:
	go build -v -ldflags "-s -w" -o dist/homekit-alsa ./cmd/homekit-alsa/main.go

.PHONY: all
all: fmt vet build
