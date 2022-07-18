.PHONY: build
build:
	go build -o out/calc .

.PHONY: run
run: build
	export PATH=$$PATH:./out