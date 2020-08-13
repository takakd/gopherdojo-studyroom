.PHONY: godoc
godoc:
	@echo "==> Starting go doc..."
	@echo "open http://localhost:8080"
	@sh -c "godoc -http=:8080"

.PHONY: build
build: fmt
	@echo "==> Building..."
	@sh -c "go build -o ./cmd/kadai1/kadai1 ./cmd/kadai1/main.go"

.PHONY: test
test:
	@echo "==> Testing..."
	@sh -c "go test -count 1 ./..."

.PHONY: fmt
fmt:
	@echo "==> Formatting..."
	@sh -c "go fmt ./..."

.PHONY: gosec
gosec:
	@echo "==> Checking security risk on code..."
	@sh -c "gosec ./..."
