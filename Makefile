build:
	go build -o build/game-of-life cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -rf build

compile: windows darwin linux wasm

windows: windows-arm64 windows-amd64

windows-arm64:
	GOOS=windows GOARCH=arm64 go build -ldflags -H=windowsgui -o build/game-of-life-arm64.exe cmd/main.go

windows-amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o build/game-of-life-amd64.exe cmd/main.go

darwin: darwin-arm64 darwin-amd64

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o build/game-of-life-darwin-arm64 cmd/main.go

darwin-amd64:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o build/game-of-life-darwin-amd64 cmd/main.go

linux: linux-arm64 linux-amd64

linux-arm64:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o build/game-of-life-linux-arm64 cmd/main.go

linux-amd64:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o build/game-of-life-linux-amd64 cmd/main.go

wasm:
	GOOS=js GOARCH=wasm go build -o build/game-of-life.wasm cmd/main.go