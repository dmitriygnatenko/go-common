package cors

import (
	"net/http"
	"strings"
)

type cors struct {
	checkOrigin  bool
	checkMethods bool
	origin       string
	methods      string
	originsList  map[string]struct{}
	methodsList  map[string]struct{}
}

func Handle(config Config, next http.Handler) http.Handler {
	instance := new(config)

	if !instance.checkOrigin && !instance.checkMethods {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			instance.handleOptions(w, r)
			w.WriteHeader(http.StatusOK)
		} else {
			instance.handleRequest(w, r)
			next.ServeHTTP(w, r)
		}
	})
}

func new(config Config) cors {
	c := cors{}

	if len(config.origin) > 0 {
		c.checkOrigin = true
		c.origin = config.origin

		origins := strings.Split(config.origin, ",")
		c.originsList = make(map[string]struct{}, len(origins))

		for i := range origins {
			c.originsList[strings.ToLower(strings.TrimSpace(origins[i]))] = struct{}{}
		}
	}

	if len(config.methods) > 0 {
		c.checkMethods = true
		c.methods = config.methods

		methods := strings.Split(config.methods, ",")
		c.methodsList = make(map[string]struct{}, len(methods))

		for i := range methods {
			c.methodsList[strings.ToUpper(strings.TrimSpace(methods[i]))] = struct{}{}
		}
	}

	return c
}

func (c cors) handleOptions(w http.ResponseWriter, r *http.Request) {
	responseHeaders := w.Header()
	origin := r.Header.Get("Origin")

	if r.Method != http.MethodOptions {
		return
	}

	responseHeaders.Add("Vary", "Origin")
	responseHeaders.Add("Vary", "Access-Control-Request-Method")
	responseHeaders.Add("Vary", "Access-Control-Request-Headers")

	if origin == "" {
		return
	}

	if !c.isOriginAllowed(origin) {
		return
	}

	if c.checkOrigin {
		responseHeaders.Set("Access-Control-Allow-Origin", c.origin)
	}

	if c.checkMethods {
		responseHeaders.Set("Access-Control-Allow-Methods", c.methods)
	}
}

func (c cors) handleRequest(w http.ResponseWriter, r *http.Request) {
	responseHeaders := w.Header()
	origin := r.Header.Get("Origin")
	responseHeaders.Add("Vary", "Origin")

	if origin == "" {
		return
	}

	if !c.isOriginAllowed(origin) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !c.isMethodAllowed(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if c.checkOrigin {
		responseHeaders.Set("Access-Control-Allow-Origin", c.origin)
	}

	if c.checkMethods {
		responseHeaders.Set("Access-Control-Allow-Methods", c.methods)
	}
}

func (c cors) isOriginAllowed(origin string) bool {
	if !c.checkOrigin {
		return true
	}

	if _, ok := c.originsList[strings.ToLower(origin)]; ok {
		return true
	}

	return false
}

func (c cors) isMethodAllowed(method string) bool {
	if !c.checkMethods {
		return true
	}

	if _, ok := c.methodsList[strings.ToUpper(method)]; ok {
		return true
	}

	return false
}
