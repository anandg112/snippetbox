FROM golang:1.22.0 AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN GOOS=linux CGO_ENABLED=0 go build -o app ./cmd/web/...

FROM alpine:latest AS run
WORKDIR /app
COPY --from=builder /build/app .
EXPOSE 4000
ENTRYPOINT ["./app"]