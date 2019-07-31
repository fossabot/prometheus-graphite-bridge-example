.PHONY: generate
generate:
	@go generate ./...

.PHONY: start
start:
	@echo "\033[0;32m» Starting Graphite on http://localhost:8080 [u/p: guest/guest]\033[0;39m"
	@docker run -d \
      --name graphite \
      -p 8080:80 \
      -p 2003:2003 \
      sitespeedio/graphite

	@echo "\033[0;32m» Open http://localhost:1337/metrics in your browser to view metrics.\033[0;39m"
	@go run examples/main.go

.PHONY: test
test:
	@go test -v --cover ./...