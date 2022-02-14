BINARY=os-api
MAIN=main.go 

run:
	go run ${MAIN}

test:
	go clean -testcache && go test ./... | grep -v 'no test files'

build: test
	go build -o ${BINARY}

.PHONY: run test build
