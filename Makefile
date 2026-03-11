.PHONY: build-hook

# Build go-hook binary and copy to all template bin/ directories
build-hook:
	cd go-hook && go build -o hook .
	mkdir -p go/bin python/bin
	cp go-hook/hook go/bin/hook
	cp go-hook/hook python/bin/hook
	rm go-hook/hook

# Run go-hook tests
test-hook:
	cd go-hook && go test -v -count=1 ./...
