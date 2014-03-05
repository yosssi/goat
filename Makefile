build:
	rm -rf ./artifacts
	mkdir ./artifacts
	mkdir ./artifacts/linux_amd64
	mkdir ./artifacts/linux_386
	mkdir ./artifacts/darwin_amd64
	mkdir ./artifacts/darwin_386
	mkdir ./artifacts/windows_amd64
	mkdir ./artifacts/windows_386
	GOOS=linux GOARCH=amd64 go build -o ./artifacts/linux_amd64/goc
	GOOS=linux GOARCH=386 go build -o ./artifacts/linux_386/goc
	GOOS=darwin GOARCH=amd64 go build -o ./artifacts/darwin_amd64/goc
	GOOS=darwin GOARCH=386 go build -o ./artifacts/darwin_386/goc
	GOOS=windows GOARCH=amd64 go build -o ./artifacts/windows_amd64/goc.exe
	GOOS=windows GOARCH=386 go build -o ./artifacts/windows_386/goc.exe
