all: build

build:
	GOOS=js GOARCH=wasm go build -o main.wasm