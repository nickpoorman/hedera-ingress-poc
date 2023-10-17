CONFIG_PATH=${HOME}/.hedera-ingress
GIT_HOST=github.com
GIT_ORG=nickpoorman
GIT_REPO=hedera-ingress
REPOSITORY=$(GIT_HOST)/$(GIT_ORG)/$(GIT_REPO)

.PHONY: init-config-dir
init-config-dir:
	@if [ ! -d "$(CONFIG_PATH)" ]; then \
		mkdir -p ${CONFIG_PATH}; \
	fi
	
.PHONY: gencert
gencert: init-config-dir
	cfssl gencert \
		-initca test/ca-csr.json | cfssljson -bare ca

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=server \
		test/server-csr.json | cfssljson -bare server

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=client \
		test/client-csr.json | cfssljson -bare client

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=client \
		-cn="root" \
		test/client-csr.json | cfssljson -bare root-client

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=client \
		-cn="nobody" \
		test/client-csr.json | cfssljson -bare nobody-client

	mv *.pem *.csr ${CONFIG_PATH}

$(CONFIG_PATH)/model.conf:
	cp test/model.conf $(CONFIG_PATH)/model.conf

$(CONFIG_PATH)/policy.csv:
	cp test/policy.csv $(CONFIG_PATH)/policy.csv

.PHONY: test
test:
test: $(CONFIG_PATH)/policy.csv $(CONFIG_PATH)/model.conf
	go test -race ./...

.PHONY: generate
generate:
	protoc rpc/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
	go generate ./...

TAG ?= 0.0.1
build-docker:
	docker build -t $(REPOSITORY):$(TAG) .

build: update-protobufs generate-hapi-methods
	CGO_ENABLED=0 go build -o ./bin/hedera-ingress ./cmd/hedera-ingress

.PHONY: run
run:
	-CGO_ENABLED=0 go run ./cmd/hedera-ingress

## Tool dependencies
TOOLS_DIR=$(PWD)/tools/bin

install-tools: install-protoc-gen-go install-protoc-gen-go-grpc

install-protoc-gen-go:
	@mkdir -p $(TOOLS_DIR)
	GOBIN=$(TOOLS_DIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	

install-protoc-gen-go-grpc:
	@mkdir -p $(TOOLS_DIR)
	GOBIN=$(TOOLS_DIR) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

clean-tools:
	@rm -rf $(TOOLS_DIR)

.PHONY: install-tools install-protoc-gen-go install-protoc-gen-go-grpc clean-tools

.PHONY: clean
clean: clean-protobufs clean-tools

## Hedera Protobuf dependencies
PROTO_REPO=https://github.com/hashgraph/hedera-protobufs.git
PROTO_DIR=tmp/hedera_protobufs
PROTO_SERVICE_DIR=$(PROTO_DIR)/services
GO_PROTOBUF_DIR=./proto
PROTO_FILES=$(wildcard $(PROTO_SERVICE_DIR)/*.proto)
# Generate M flags for each proto file
M_FLAGS=$(foreach file, $(PROTO_FILES), --go_opt=M$(notdir $(file))=github.com/hashgraph/hedera-protobufs-go/services/proto --go-grpc_opt=M$(notdir $(file))=github.com/hashgraph/hedera-protobufs-go/services/proto)

update-protobufs: install-tools clone-proto-repo generate-go-protos

# Clones the protobuf repository to a temporary directory.
clone-proto-repo:
	@if [ ! -d "$(PROTO_SERVICE_DIR)" ]; then \
		git clone $(PROTO_REPO) $(PROTO_DIR); \
	else \
		echo "Directory $(PROTO_SERVICE_DIR) already exists, skipping clone."; \
	fi

# Generates Go proto files from the protobufs.
generate-go-protos: clone-proto-repo
	@mkdir -p $(GO_PROTOBUF_DIR)
	protoc -I=$(PROTO_SERVICE_DIR) $(PROTO_SERVICE_DIR)/*.proto --go_out=$(GO_PROTOBUF_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GO_PROTOBUF_DIR) --go-grpc_opt=paths=source_relative $(M_FLAGS)

clean-protobufs:
	@rm -rf $(PROTO_DIR) $(GO_PROTOBUF_DIR)

.PHONY: clone-proto-repo generate-go-protos clean-protobufs

# Generate the hapi methods given the generated protobufs
.PHONY: generate-hapi-methods
generate-hapi-methods:
	go run ./cmd/hapi-methods-gen/main.go -dir proto/ > internal/server/hapi/hapi.gen.go

## Hedera Go SDK test dependencies
HEDERA_GO_SDK_REPO=https://github.com/hashgraph/hedera-sdk-go.git
HEDERA_GO_SDK_DIR=tmp/hedera-sdk-go
HEDERA_GO_SDK_CONFIG=$(shell pwd)/test/hedera-go-client-config.json
HEDERA_GO_SDK_LOG_LEVEL=ERROR # ERROR, WARNING, INFO, DEBUG, TRACE
HEDERA_GO_SDK_SPECIFIC_TEST=-run ^TestIntegrationTopic # -run ^TestIntegrationAccountCreateTransactionCanExecute$

download-hedera-go-sdk-e2e-tests:
	@if [ ! -d "$(HEDERA_GO_SDK_DIR)" ]; then \
		git clone $(HEDERA_GO_SDK_REPO) $(HEDERA_GO_SDK_DIR); \
	else \
		echo "Directory $(HEDERA_GO_SDK_DIR) already exists, skipping clone."; \
	fi

test-e2e: download-hedera-go-sdk-e2e-tests
	cd $(HEDERA_GO_SDK_DIR) && HEDERA_SDK_GO_LOG_LEVEL="$(HEDERA_GO_SDK_LOG_LEVEL)" CGO_ENABLED=0 CONFIG_FILE=$(HEDERA_GO_SDK_CONFIG) go test -tags="e2e" -v -timeout 30s $(HEDERA_GO_SDK_SPECIFIC_TEST)

.PHONY: run-hedera-go-sdk-e2e-tests

## Hedera local-node test dependencies
HEDERA_LOCAL_NODE_DIR=tmp/hedera-local-node
HEDERA_LOCAL_NODE_VERSION=0.13.0
HEDERA_LOCAL_NODE_REPO=https://github.com/hashgraph/hedera-local-node.git

download-local-node:
	@if [ ! -d "$(HEDERA_LOCAL_NODE_DIR)" ]; then \
		git clone $(HEDERA_LOCAL_NODE_REPO) $(HEDERA_LOCAL_NODE_DIR); \
	else \
		echo "Directory $(HEDERA_LOCAL_NODE_DIR) already exists, skipping clone."; \
	fi

run-local-node: download-local-node
	cd $(HEDERA_LOCAL_NODE_DIR) && docker compose up

.PHONY: download-local-node run-local-node

## Hedera mirror-node http client
# Variables
SPEC_PATH = tmp/mirror-node-spec.yml
MIRROR_NODE_CLIENT_OUTPUT_DIR = tmp/mirror-node-client
OPEN_API_GENERATOR = openapi-generator-cli
OPEN_API_GENERATOR_CONFIG_PATH = test/openapi-generator.config
SPEC_URL = https://mainnet-public.mirrornode.hedera.com/api/v1/docs/openapi.yml
MIRROR_NODE_CLIENT_DIR = mirrorclient

# Fetch the OpenAPI spec if it doesn't exist.
fetch-mirror-node-spec-if-not-exists:
	@if [ ! -f $(SPEC_PATH) ]; then \
		echo "Spec file not found. Fetching from $(SPEC_URL)..."; \
		curl -o $(SPEC_PATH) $(SPEC_URL); \
	fi

# Check if openapi-generator-cli is installed, if not install it.
install-openapi-generator-cli:
	@if ! command -v $(OPEN_API_GENERATOR) &> /dev/null; then \
		echo "openapi-generator-cli is not installed, installing via npm..."; \
		npm install @openapitools/openapi-generator-cli -g; \
	else \
		echo "openapi-generator-cli is already installed"; \
	fi

# Generate the Go client.
generate-mirror-node-client: fetch-mirror-node-spec-if-not-exists install-openapi-generator-cli
	$(OPEN_API_GENERATOR) generate -i $(SPEC_PATH) --package-name mirrorclient --api-package mirrorclient --git-host $(GIT_HOST) --git-user-id $(GIT_ORG) --git-repo-id $(GIT_REPO)/mirrorclient -g go -o $(MIRROR_NODE_CLIENT_OUTPUT_DIR) --config $(OPEN_API_GENERATOR_CONFIG_PATH)
	rm -rf $(MIRROR_NODE_CLIENT_DIR)
	mv $(MIRROR_NODE_CLIENT_OUTPUT_DIR) $(MIRROR_NODE_CLIENT_DIR)

.PHONY: generate-mirror-node-client install-openapi-generator-cli fetch-mirror-node-spec-if-not-exists