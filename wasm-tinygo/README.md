# WebAssemble with Go

## Commands

```bash

# Build go script for WASM
> $env:GOOS='js'; $env:GOARCH='wasm'; go build -o public/wasm.wasm

# Build go calculator script for wasm
> $env:GOOS='js'; $env:GOARCH='wasm'; go build -o public/calculator.wasm ./wasm/main.go

# Run go script 
> $env:GOOS='windows'; $env:GOARCH='amd64'; go run main.go
```

## VSCode Settings for WASM

```json
"go.toolsEnvVars": {
  "GOOS": "js",
  "GOARCH": "wasm",
},
```