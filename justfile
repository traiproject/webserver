set shell := ["bash", "-c"]

default:
	@just --list

[parallel]
dev: dev-templ dev-js dev-css dev-go

dev-templ:
	templ generate -watch

dev-js:
	bun run watch:js

dev-css:
	bun run watch:css 

dev-go:
	air

build: build-templ build-js build-css build-go

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
	fd "_templ.go" -X rm
