SHORT_NAME = simple-http-server
IMAGE = cainelli/simple-http-server
DEV_TAG = 1.0.0-dev
GOOS ?= linux
GOARCH ?= amd64
REV = $(shell git rev-parse --short HEAD)
VERSION = $(shell echo $(BUILD_VERSION))
CURRENTPATH = $(shell echo $(PWD))
HOMEDIR = $(shell echo $(HOME))


release: image-build image-push

image-build:
	docker build --build-arg BUILD_VERSION="1.0.0-dev" . -f Dockerfile -t ${IMAGE}:${DEV_TAG}

image-push:
	docker push ${IMAGE}:${DEV_TAG}

up:
	go run main.go

build:
	go build -o simple-http-server cmd/main.go
