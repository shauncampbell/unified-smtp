clean:
	rm -rf smtp.*
lint:
	golangci-lint run ./internal/... ./cmd/... ./pkg/...
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o smtp.linux_amd64 ./cmd/smtp
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o smtp.darwin_amd64 ./cmd/smtp
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o smtp.windows_amd64.exe ./cmd/smtp

docker:
	docker build -f ./cmd/smtp/Dockerfile -t shauncampbell/smtp:local .