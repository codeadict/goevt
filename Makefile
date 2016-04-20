.PHONY: test run update format install

install:
	go get github.com/stretchr/testify/assert
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get -v
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs go get -v

update:
	go get -u all

format:
	gofmt -l -w -s .
	go fix ./...

test:
	@rm -f /tmp/goevt-coverprofile.out
	@go test -v -coverprofile=/tmp/goevt-coverprofile.out
	@go tool cover -func=/tmp/goevt-coverprofile.out
	@rm -f /tmp/goevt-coverprofile.out
	exit `gofmt -l -s -e . | wc -l`
