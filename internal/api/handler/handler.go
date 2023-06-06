package handler

import (
	"net/http"

	"github.com/larikhide/website-monitor/internal/app/website"
)

type Router struct {
	*http.ServeMux
	ws *website.Websites
}

func NewRouter(ws *website.Websites) *Router {
	r := &Router{
		ServeMux: http.NewServeMux(),
		ws:       ws,
	}
	r.Handle("/getAccessTime", r.AuthMiddleware(http.HandlerFunc(r.GetAccessTime)))
	r.Handle("/GetMinAccessURL", r.AuthMiddleware(http.HandlerFunc(r.GetMinAccessURL)))
	r.Handle("/GetMinAccessURL", r.AuthMiddleware(http.HandlerFunc(r.GetMaxAccessURL)))
	return r
}

func (rt *Router) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if u, p, ok := r.BasicAuth(); !ok || !(u == "admin" && p == "admin") {
				http.Error(w, "unautorized", http.StatusUnauthorized)
				return
			}
			// r = r.WithContext(context.WithValue(r.Context(), 1, 0))
			next.ServeHTTP(w, r)
		},
	)
}
func (rt *Router) GetAccessTime(w http.ResponseWriter, r *http.Request)   {}
func (rt *Router) GetMinAccessURL(w http.ResponseWriter, r *http.Request) {}
func (rt *Router) GetMaxAccessURL(w http.ResponseWriter, r *http.Request) {}
