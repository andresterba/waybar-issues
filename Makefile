GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
BINARY_NAME=waybar-issues

all: build
run: build
	./$(BINARY_NAME)
build: 
	$(GOBUILD) -o $(BINARY_NAME)

