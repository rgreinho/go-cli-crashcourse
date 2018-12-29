# Project name.
PROJECT_NAME = crashcourse

# Makefile parameters.
TAG ?= $(shell git describe)

# General.
SHELL = /bin/bash
TOPDIR = $(shell git rev-parse --show-toplevel)

# Project specifics.
BUILD_DIR = dist
PLATFORMS = linux darwin
OS = $(word 1, $@)
GOOS = $(shell uname -s | tr A-Z a-z)
GOARCH = amd64

default: build

help: # Display help
	@awk -F ':|##' \
		'/^[^\t].+?:.*?##/ {
			printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
		}' $(MAKEFILE_LIST) | sort

bootstrap-osx: ## Bootstrap a Go setup for OSX
	@bash scripts/bootstrap-osx.sh
	@bash scripts/bootstrap-shell.sh

build: ## Build the project for the current platform
	mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(PROJECT_NAME)-$(TAG)-$(GOOS)-$(GOARCH)

ci: ci-linters ci-tests ## Run all the CI targets

ci-linters: ## Run the static analyzers
	gometalinter --skip=vendor ./...

ci-tests: ## Run the unit tests
	go test pkg/*

clean: clean-code ## Clean everything (!DESTRUCTIVE!)

clean-code: ## Remove unwanted files in this project (!DESTRUCTIVE!)
	@cd $(TOPDIR) && git clean -ffdx && git reset --hard

dist: $(PLATFORMS) ## Package the project for all available platforms

setup: ## Setup the full environment (default)
	dep ensure
	gometalinter --install
	go get -u github.com/spf13/cobra/cobra
	go get -u github.com/derekparker/delve/cmd/dlv

.PHONY: help bootstrap-osx build ci ci-linters ci-tests clean clean-code dist setup

$(PLATFORMS): # Build the project for all available platforms
	mkdir -p $(BUILD_DIR)
	GOOS=$(OS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(PROJECT_NAME)-$(TAG)-$(OS)-$(GOARCH)
.PHONY: $(PLATFORMS)
