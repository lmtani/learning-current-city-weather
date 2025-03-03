tidy:
	go mod tidy
	gofmt -w .
	goimports -w .


test:
	go test -v ./...
