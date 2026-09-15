.PHONY: run generate validate build tidy test

run:
	go run ./cmd/gophoner $(ARGS)

generate:
	go generate ./...

validate: generate
	deadcode ./...

build: validate
	go build -o gophoner ./cmd/gophoner

tidy:
	go mod tidy

test:
	go test ./internal/report/... -run TestPrintPreview -v