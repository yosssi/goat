build:
	rm -rf ./artifacts
	mkdir ./artifacts
	GOOS=linux GOARCH=amd64 go build -o ./artifacts/goat-linux-amd64
	GOOS=linux GOARCH=386 go build -o ./artifacts/goat-linux-386
	GOOS=darwin GOARCH=amd64 go build -o ./artifacts/goat-darwin-amd64
	GOOS=darwin GOARCH=386 go build -o ./artifacts/goat-darwin-386
	GOOS=windows GOARCH=amd64 go build -o ./artifacts/goat-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -o ./artifacts/goat-windows-386.exe
