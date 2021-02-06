GO=go
MAKE=make
API_ROOT=https://www.dnd5eapi.co/api

integration-test:
	API_ROOT=$(API_ROOT) $(GO) test -count 1 -tags "unit integration" ./...