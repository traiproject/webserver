package router

import (
	"io/fs"
	"net/http"

	"example.com/webserver/internal/app/config"
	"example.com/webserver/internal/infrastructure/http/middleware"
	"example.com/webserver/web"
)

func serveStatic(mux *http.ServeMux, cfg *config.Config) {
	staticCache := middleware.CacheStatic(cfg.IsProduction())

	staticFS, err := fs.Sub(web.StaticFiles, "static")
	if err != nil {
		panic(err)
	}

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.FS(staticFS)))
	mux.HandleFunc("GET /static/", staticCache(staticHandler.ServeHTTP))
}
