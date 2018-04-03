.PHONY: run test image

VERSION := 0.1.0

run:
	go run ./echo.go

image:
	docker build -t echo:${VERSION} .

test:
	go test ./...