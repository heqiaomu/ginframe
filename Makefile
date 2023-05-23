PROJECT := ginframe
# ==============================================================================
# Build Options

ROOT_PACKAGE=github.com/Heqiaomu/$(PROJECT)
# version通用包路径
VERSION_PACKAGE=$(ROOT_PACKAGE)/pkg/version

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
OS := $(shell uname)
ARCH := $(shell go env GOARCH)

ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR) && pwd -P))
endif

ifeq ($(origin OUTPUT_PREFIX),undefined)
OUTPUT_PREFIX := output
endif

ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ROOT_DIR)/$(OUTPUT_PREFIX)
endif

# blocface应用名前缀
GO_APP_PRE := Heqiaomu-

# need git repo info
GO_LDFLAGS += -X $(VERSION_PACKAGE).gitTag=$(GIT_TAG) -X $(VERSION_PACKAGE).version=$(VERSION) -X $(VERSION_PACKAGE).buildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ') -X $(VERSION_PACKAGE).gitCommit=$(GIT_COMMIT) -X $(VERSION_PACKAGE).gitTreeState=$(GIT_TREE_STATE) -X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn -s -w

ifeq ($(GOOS),windows)
	GO_OUT_EXT := .exe
endif

ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS 2>/dev/null)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH 2>/dev/null)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
endif

.PHONY: clear
clear:
	@go clean

.PHONY: run
run:
	@cd cmd/ && go run main.go --mode=debug

.PHONY: linux.%
linux.%: ## build.amd64 编译linux版本的amd或者arm
	$(eval bin := $(word 1,$(subst ., ,$*)))
	@echo "===========> Building binary for GOOS=linux GOARCH=$(bin)"
	@mkdir -p $(OUTPUT_DIR)/linux/$(bin)
	@CGO_ENABLED=0 GOOS=linux GOARCH=$(bin) go build -o $(OUTPUT_DIR)/linux/$(bin)/$(GO_APP_PRE)$(PROJECT)$(GO_OUT_EXT) -ldflags "$(GO_LDFLAGS)" $(ROOT_DIR)/cmd/$(PROJECT)
