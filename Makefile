all: win lin

name := probelistener

win:
	GOOS=windows GOARCH=amd64 go build -o out/main.exe -trimpath -ldflags="-s -w" ./cmd/
	upx64 out/main.exe -f -o out/$(name).exe
	rm -rf out/main.exe

lin:
	GOOS=linux GOARCH=amd64 go build -o out/main -trimpath -ldflags="-s -w" ./cmd/
	upx64 out/main -f -o out/$(name)
	rm -rf out/main