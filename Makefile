.PHONY: run
run: 
	go run --race ./cmd/apiserver/

.PHONY: build
build:
	go build -o openmetrics ./cmd/apiserver/

.PHONY: test
test:
	go test --race --cover -v ./...

.PHONY: test-html
test-html:
	go test --race --coverprofile=c.out ./...	
	go tool cover -html=c.out

.DEFAULT_GOAL=run