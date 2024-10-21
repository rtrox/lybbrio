package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
)

const VITE_DEV = "http://localhost:5173"

func WebRoutes(isDev bool, proxyPath string, assetFolder string) http.Handler {
	r := chi.NewRouter()
	if isDev {
		r.Get("/*", FrontendProxy(proxyPath))
	} else {
		r.Get("/*", FrontentDist(assetFolder))
	}
	return r
}

func FrontendProxy(proxyPath string) http.HandlerFunc {
	remote, err := url.Parse(proxyPath)
	if err != nil {
		// Should be unreachable
		panic(err)
	}
	proxy := httputil.ReverseProxy{Director: func(r *http.Request) {
		r.URL.Scheme = remote.Scheme
		r.URL.Host = remote.Host
		r.URL.Path = remote.Path + r.URL.Path
		r.Host = remote.Host
	}}
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func FrontentDist(assetFolder string) http.HandlerFunc {
	assetDir := http.Dir(assetFolder)
	return func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(assetDir))
		if r.URL.Path != "/" {
			fullPath, err := url.JoinPath(assetFolder, strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
			if err != nil {
				statusCodeResponse(w, r, http.StatusNotFound)
				return
			}
			_, err = os.Stat(fullPath)
			if err != nil {
				if !os.IsNotExist(err) {
					statusCodeResponse(w, r, http.StatusInternalServerError)
					return
				}
				r.URL.Path = "/"
			}
		}
		fs.ServeHTTP(w, r)
	}
}
