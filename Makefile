# Go commands
GOCMD 	= go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST 	= $(GOCMD) test
GOGET 	= $(GOCMD) get

# Filepaths
BUILD_FOLDER	= build
BINARY_NAME 	= $(BUILD_FOLDER)/gotv
COVERAGE_OUT 	= $(BUILD_FOLDER)/coverage.out
COVERAGE_HTML 	= $(BUILD_FOLDER)/coverage.html


# Default target
default: clean build run

# Build target
build:
	@$(GOBUILD) -o $(BINARY_NAME)

# Built and run target
run: build
	@./$(BINARY_NAME)

# Clean target
clean:
	@$(GOCLEAN)
	@rm -rf $(BUILD_FOLDER)

# Test target
test:
	@mkdir -p $(BUILD_FOLDER)
	@$(GOTEST) -v -coverprofile=$(COVERAGE_OUT) ./...
	@go tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

# Install dependencies
deps:
	@$(GOGET) -v ./...
