GO=go
MAKE=make
API_ROOT=http://localhost:3000/api
FLAGS=

integration-test:
	API_ROOT=$(API_ROOT) $(GO) test -count 1 -tags "unit integration" $(FLAGS) ./...

generate:
	$(GO) generate $(FLAGS) ./...

run-examples:
	$(GO) run ./examples/...