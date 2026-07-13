GO_FILES := $(shell find . -type f -name '*.go' -not -path './.git/*')

.PHONY: run fmt fmt-check test race vet check

run:
	go run .

fmt:
	gofmt -w $(GO_FILES)

fmt-check:
	@unformatted="$$(gofmt -l $(GO_FILES))"; \
	if [ -n "$$unformatted" ]; then \
		printf '%s\n' "$$unformatted"; \
		exit 1; \
	fi

test:
	go test ./...

race:
	go test -race ./...

vet:
	go vet ./...

check: fmt-check vet test race
