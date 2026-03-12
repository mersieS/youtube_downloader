APP_NAME := ytmp3
VERSION  ?= 0.1.0
LDFLAGS  := -s -w -X main.version=$(VERSION)
BUILD_DIR := build

.PHONY: build build-all clean install release

build:
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME) .

build-darwin-arm64:
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 .

build-darwin-amd64:
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 .

build-windows-amd64:
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe .

build-all: build-darwin-arm64 build-darwin-amd64 build-windows-amd64

release: build-all
	@cd $(BUILD_DIR) && tar czf $(APP_NAME)-darwin-arm64.tar.gz $(APP_NAME)-darwin-arm64
	@cd $(BUILD_DIR) && tar czf $(APP_NAME)-darwin-amd64.tar.gz $(APP_NAME)-darwin-amd64
	@cd $(BUILD_DIR) && zip -q $(APP_NAME)-windows-amd64.zip $(APP_NAME)-windows-amd64.exe
	@cd $(BUILD_DIR) && shasum -a 256 *.tar.gz *.zip > checksums.txt
	@echo "Release artifacts:"
	@ls -lh $(BUILD_DIR)/*.tar.gz $(BUILD_DIR)/*.zip
	@echo ""
	@cat $(BUILD_DIR)/checksums.txt

install: build
	cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "$(APP_NAME) installed to /usr/local/bin/$(APP_NAME)"

clean:
	rm -rf $(BUILD_DIR)
