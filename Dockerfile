FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM scratch AS service_a

WORKDIR /app
COPY --from=builder /app/bin/svc-a .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./svc-a"]


FROM scratch AS service_b

WORKDIR /app
COPY --from=builder /app/bin/svc-b .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./svc-b"]