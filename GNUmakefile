default: testacc

# Run acceptance tests
.PHONY: testacc testcover
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -v -coverpkg=all -timeout 120m

testcover:
	# go test -v -timeout=5s ./src/...
	# go test --cover --short --timeout 5s ./src/...
	TF_ACC=1 gotestsum --junitfile report.xml --format testname --raw-command go test --cover --timeout 120s --tags musl  --coverprofile=coverage.txt --covermode=atomic --coverpkg "$(shell go list ./...  | tr '\n' ",")" --json ./... 
	go tool cover -func coverage.txt
	go-cover-treemap -statements -w 1900 -h 1200 -coverprofile coverage.txt  > coverage.svg
	go run github.com/boumenot/gocover-cobertura < coverage.txt > coverage.xml

test:

install:
	go install -v ./...
	module=$(shell grep module go.mod | head -1 | cut -d ' ' -f 2)

test:
	TF_ACC=1 go test -v --timeout 120s ./... 

generate:
	go generate

upgrade:
	go install github.com/oligot/go-mod-upgrade@latest
	go-mod-upgrade