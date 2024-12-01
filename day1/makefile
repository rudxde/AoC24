install: go.mod go.sum
	@echo installing dependencies
	go mod download
	@echo done installing dependencies
tidy:
	go mod tidy
.PHONY: build install
build:
	go build -o dist/aoc_day1 ./
	chmod +x dist/aoc_day1
	@echo build completed

start: build
	./dist/aoc_day1
