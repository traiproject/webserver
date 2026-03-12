// Package web contains the embedded fs for the static files for the webserver.
package web

import "embed"

// StaticFiles contains the static files for the webserver.
//
//go:embed static/*
var StaticFiles embed.FS
