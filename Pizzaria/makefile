PROJECT_NAME=pizzaria

run: 
	go run ./cmd/main.go

build: binaries
	GOOS=darwin GOARCH=amd64 go build -o ./binaries/${PROJECT_NAME}-darwin ./cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o ./binaries/${PROJECT_NAME}-linux ./cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o ./binaries/${PROJECT_NAME}-windows.exe ./cmd/main.go

clean:
	go clean
	rm -f ./binaries/${PROJECT_NAME}-darwin
	rm -f ./binaries/${PROJECT_NAME}-linux
	rm -f ./binaries/${PROJECT_NAME}-windows.exe

binaries:
	mkdir -p binaries