bin: 
	mkdir -p

bin/codetool: bin $(shell find . -name '*.go')
	go build -o bin/codetool .

.PHONY: build
build: bin/codetool
	@echo Build complete.

.PHONY: clean
clean:
	rm -rf build

.PHONY: run
run-analyze: build
	bin/codetool codebuild analyze