build:
	@go build -o bin/posts

run: build
	@./bin/posts

dev:
	@gin run *go