FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM scratch

WORKDIR /app
COPY --from=builder /app/bin/server .
# Copy CA certificates from the builder image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./server"]
