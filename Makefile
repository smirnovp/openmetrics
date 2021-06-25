.PHONY: run
run: 
	go run --race ./cmd/apiserver/

.DEFAULT_GOAL=run