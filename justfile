set shell := ["bash", "-c"]

default:
	@just --list

dev:
	air

build: build-prepare build-go

clean:
	rm -rf ./tmp
	rm -rf ./web/static/js/*
	rm -rf ./web/static/css/*
	fd -I -t f -g "*_templ.go" -X rm
	fd -I -t f -g "*.sql.go" -X rm

[parallel]
build-prepare: build-templ build-js build-css build-db

build-templ:
	templ generate

build-js:
	bun run build:js

build-css:
	bun run build:css

build-db:
	sqlc generate

build-go:
	go build -o ./tmp/main ./cmd/web

