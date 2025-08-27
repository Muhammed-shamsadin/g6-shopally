package router

import (
	"net/http"
	"strings"

	apphandler "github.com/shopally-ai/internal/adapter/handler"
)

// Deps contains all handlers that the router should mount.
type Deps struct {
	FX *apphandler.FXHandler
}

// Options control router behavior like base path and middlewares.
type Options struct {
	BasePath    string
	Middlewares []func(http.Handler) http.Handler
}

// Build constructs the http.Handler with all routes mounted and optional middlewares applied.
// Example: handler := router.Build(router.Deps{FX: fxHandler}, router.Options{BasePath: "/api/v1"})
func Build(d Deps, opts ...Options) http.Handler {
	mux := http.NewServeMux()

	var base string
	var mws []func(http.Handler) http.Handler
	if len(opts) > 0 {
		base = strings.TrimRight(opts[0].BasePath, "/")
		mws = opts[0].Middlewares
	}

	// Mount feature routes
	mountFX(mux, d.FX, base)

	// Wrap with middlewares (outermost first)
	var h http.Handler = mux
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func mountFX(mux *http.ServeMux, fx *apphandler.FXHandler, base string) {
	if fx == nil {
		return
	}
	path := "/fx"
	if base != "" {
		path = base + path
	}
	mux.HandleFunc(path, fx.GetFX)
}
