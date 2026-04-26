FROM --platform=linux/amd64 golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/bin/service .

FROM --platform=linux/amd64 gcr.io/distroless/base-debian12:nonroot AS final
WORKDIR /app
COPY --from=builder /app/bin/service .
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app/service"]