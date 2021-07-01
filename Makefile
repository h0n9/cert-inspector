BIN=cert-inspector
PKG=./cmd/cert-inspector

build:
	go build -o $(BIN) $(PKG)

run:
	go run $(PKG)