FROM golang:1.9.5-alpine AS builder
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build echo.go

FROM alpine
WORKDIR /app

COPY --from=builder /app/echo ./

CMD ./echo