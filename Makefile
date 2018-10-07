  # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
DIST=dist
GOOS=linux
GOARCH=amd64
BINARY_NAME=ja-bot
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build dist
dist:
	[ -d $(DIST) ] || mkdir $(DIST)
	cp *config.yaml $(DIST)/
	mv $(BINARY_NAME)* $(DIST)/
build: 
		# env GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o $(DIST)/$(BINARY_NAME)-$(GOOS)-$(GOARCH) -v
		# env GOOS=darwin GOARCH=$(GOARCH) $(GOBUILD) -o $(DIST)/$(BINARY_NAME)-darwin-$(GOARCH) -v
		gox -os="linux darwin" -arch="amd64" -verbose ./...
		
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)*
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/markbates/goth
		$(GOGET) github.com/markbates/pop


# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v