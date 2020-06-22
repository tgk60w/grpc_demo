default: dev

all: generate
	#GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o app-windows-x64.exe -ldflags "-s -w" app.go
#	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o app-windows-386.exe -ldflags "-s -w" app.go
	#GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app-linux-x64 -ldflags "-s -w" app.go
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o app-linux-386 -ldflags "-s -w" app.go
	#GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o app-darwin-x64 -ldflags "-s -w" app.go

generate:
	go generate

update:
	go list -u -m all

dev: generate
	go run app.go --module=vendor

golint:
	#go get -u golang.org/x/lint/golint
	golint ./...
#	find . -type d | xargs -L 1 golint
gofmt:
	gofmt -s -l ./

protoc:
	protoc --proto_path=./protos --go_out=plugins=grpc:./protos ./protos/transfer.proto