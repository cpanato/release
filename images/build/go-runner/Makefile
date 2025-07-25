# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# include the common image-building Makefiles
include $(CURDIR)/../../Makefile.common-image

IMGNAME = go-runner
APP_VERSION = $(shell cat VERSION)
GO_MAJOR_VERSION ?= 1.24
REVISION ?= 0
GO_VERSION ?= 1.24.5
OS_CODENAME ?= bookworm

# Build args
DISTROLESS_REGISTRY ?= gcr.io/distroless
DISTROLESS_IMAGE ?= static-debian12
BUILDER_IMAGE ?= golang:$(GO_VERSION)-$(OS_CODENAME)

# Configuration
CONFIG = go$(GO_MAJOR_VERSION)-$(OS_CODENAME)
BUILD_METADATA = go$(GO_VERSION)-$(OS_CODENAME).$(REVISION)

IMAGE_VERSION = $(APP_VERSION)-$(BUILD_METADATA)

PLATFORMS ?= linux/amd64 linux/arm64 linux/arm linux/ppc64le linux/s390x

HOST_GOOS ?= $(shell go env GOOS)
HOST_GOARCH ?= $(shell go env GOARCH)
GO_BUILD ?= go build

.PHONY: all build clean

.PHONY: all
all: build

.PHONY: build
build:
	$(GO_BUILD)

.PHONY: clean
clean:
	rm go-runner

BUILD_ARGS = --build-arg=BUILDER_IMAGE=$(BUILDER_IMAGE) \
             --build-arg=DISTROLESS_IMAGE=$(DISTROLESS_REGISTRY)/$(DISTROLESS_IMAGE)
