.PHONY: test

test:
	@rm -f /tmp/goevt-coverprofile.out
	@go test -v -coverprofile=/tmp/goevt-coverprofile.out
	@go tool cover -func=/tmp/goevt-coverprofile.out
	@rm -f /tmp/goevt-coverprofile.out
