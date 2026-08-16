OWNER ?=
BINARY := terraform-provider-stalwart

default: build

build:
	go build -v ./...

install:
	go install -v ./...

test:
	go test -v -cover -timeout=120s ./...

testacc:
	TF_ACC=1 go test -v -cover -timeout=120m ./...

lint:
	golangci-lint run

fmt:
	gofmt -s -w -e .

docs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name stalwart

schema:
	@test -n "$(STALWART_URL)" || (echo "STALWART_URL is required"; exit 1)
	@test -n "$(STALWART_VERSION)" || (echo "STALWART_VERSION is required"; exit 1)
	curl -sSf -u "$(STALWART_USER):$(STALWART_PASSWORD)" "$(STALWART_URL)/api/schema" \
		-o schema/$(STALWART_VERSION).json
	@echo "wrote schema/$(STALWART_VERSION).json"

rename:
	@test -n "$(OWNER)" || (echo "usage: make rename OWNER=your-github-account"; exit 1)
	grep -rl 'OWNER' --include='*.go' --include='go.mod' --include='*.md' --include='*.tf' . \
		| xargs sed -i '' 's|github.com/OWNER/|github.com/$(OWNER)/|g; s|registry.terraform.io/OWNER/|registry.terraform.io/$(OWNER)/|g; s|"OWNER/stalwart"|"$(OWNER)/stalwart"|g'
	go mod tidy
	@echo "module path now targets $(OWNER)"

.PHONY: build install test testacc lint fmt docs schema rename
