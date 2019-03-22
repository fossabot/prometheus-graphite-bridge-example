.PHONY: generate
generate:
	@go generate ./...

.PHONY: start
start:
	@echo "Open http://localhost:1337/metrics in your browser to view metrics."
	@go run examples/main.go

.PHONY: test
test:
	@go test -v --cover ./...