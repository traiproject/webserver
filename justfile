set shell := ["bash", "-c"]

default:
	@just --list

dev:
	air

[parallel]
build-prepare: build-templ build-js build-css

build: build-prepare build-go

build-templ:
	templ generate

build-js:
	bun run build:js

build-css:
	bun run build:css

build-go:
	go build -o ./tmp/main ./cmd/web

clean:
	rm -rf ./tmp
	rm -rf ./web/static/js/*
	rm -rf ./web/static/css/*
	fd -I -t f -g "*_templ.go" -X rm
