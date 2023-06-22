package openapiview

import (
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

type Middleware struct {
	pathPrefix  string
	openapiData []byte
}

func (m *Middleware) Process(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, m.pathPrefix) {
			path := strings.TrimPrefix(r.URL.Path, m.pathPrefix)
			path = "static" + path

			if path == "static/openapi" {
				_, _ = w.Write(m.openapiData)
				return
			}

			data, err := fs.ReadFile(path)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			contentType := mime.TypeByExtension(filepath.Ext(r.URL.Path))
			w.Header().Set("Content-Type", contentType)
			_, _ = w.Write(data)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewMiddleware(pathPrefix string, openapiData []byte) *Middleware {
	return &Middleware{pathPrefix: pathPrefix, openapiData: openapiData}
}
