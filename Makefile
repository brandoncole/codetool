.PHONY: build clean tests sample

bin: 
	mkdir -p

bin/codetool: bin $(shell find . -name '*.go')
	go build -o bin/codetool .

build: bin/codetool
	@echo Build complete.