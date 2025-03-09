tidy:
	go mod tidy
	gofmt -w .
	goimports -w .

test:
	go test -v ./...

build-service-a:
	mkdir -p bin
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ./bin/svc-a ./cmd/svc_a/main.go

build-service-b:
	mkdir -p bin
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ./bin/svc-b ./cmd/svc_b/main.go

build:
	make build-service-a
	make build-service-b
