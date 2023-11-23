package handlers

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type ctxKey struct{}

var routes = []route{
	newRoute("GET", "/", RootHandler),
	newRoute("GET", "/login/github", githubLoginHandler),
	newRoute("GET", "/login/github/callback", githubCallbackHandler),
	newRoute("GET", "/loggedin", func(w http.ResponseWriter, r *http.Request) {
		LoggedinHandler(w, r, "")
	}),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
    log.Printf("Url path: %s\n", r.URL.Path)
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ","))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
    
	http.NotFound(w, r)
}
