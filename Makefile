GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOFORMAT=$(GOCMD) fmt
BINARY_NAME=waybar-issues

all: build
run: build
	./bazel-bin/waybar-issues_/waybar-issues
build: 
	bazel build //... 
format:
	$(GOFORMAT) ./...
test:
	bazel test //... --test_output=errors
