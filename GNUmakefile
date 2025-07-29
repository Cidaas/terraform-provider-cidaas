TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=Cidaas
NAME=cidaas
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=linux_amd64
GO_LINT ?= golangci-lint
GO_LINT_CONFIG_PATH ?= ./lint.yml
TEST?=./internal/...

default: install

build:
	go build -o ${BINARY}

deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o ./bin/${BINARY}_${VERSION}_darwin_arm64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

# test:
# 	go test -i $(TEST) || exit 1
# 	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./internal

fmtcheck: deps
	@./scripts/gofmtcheck.sh
	$(GO_LINT) run --config $(GO_LINT_CONFIG_PATH) $(GO_LINT_ARGS)

test: fmtcheck
	go test $(TEST) $(TESTARGS) -timeout=5m -parallel=4

test-ci:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -parallel=1 -coverprofile .coverage.txt ./...
	go tool cover -func .coverage.txt
	go tool cover -html=.coverage.txt -o coverage.html