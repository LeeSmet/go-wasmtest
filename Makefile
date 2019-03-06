all: build

build:
	GOOS=js GOARCH=wasm go build -o main.wasm

jshelper:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .

.PHONY: build
