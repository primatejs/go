GOOS=js GOARCH=wasm go build -o route.wasm route.go request.go
node run.js
