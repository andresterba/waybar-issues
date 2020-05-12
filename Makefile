GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=waybar-issues

all: build
run: build
	./$(BINARY_NAME)
build: 
	$(GOBUILD) -o $(BINARY_NAME)
test:
	$(GOTEST) ./...
