tidy:
	go mod tidy
	gofmt -w .
	goimports -w .

test:
	go test -v ./...


build:
	mkdir -p bin
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ./bin/server ./cmd/server/main.go
