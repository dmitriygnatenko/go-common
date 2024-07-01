package cors

import (
	"net/http"
	"strings"
)

type CORS struct {
	checkOrigin  bool
	checkMethods bool
	origin       string
	methods      string
	originsList  map[string]struct{}
	methodsList  map[string]struct{}
}

func NewCORSMiddleware(c Config) (*CORS, error) {
	cors := CORS{}

	if len(c.origin) > 0 {
		cors.checkOrigin = true
		cors.origin = c.origin

		origins := strings.Split(cors.origin, ",")
		cors.originsList = make(map[string]struct{}, len(origins))

		for i := range origins {
			cors.originsList[strings.ToLower(origins[i])] = struct{}{}
		}
	}

	if len(c.methods) > 0 {
		cors.checkMethods = true
		cors.methods = c.methods

		methods := strings.Split(cors.methods, ",")
		cors.methodsList = make(map[string]struct{}, len(methods))

		for i := range methods {
			cors.methodsList[strings.ToUpper(methods[i])] = struct{}{}
		}
	}

	return &cors, nil
}

func (c CORS) Handle(next http.Handler) http.Handler {
	if !c.checkOrigin && !c.checkMethods {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			c.handleOptions(w, r)
			w.WriteHeader(http.StatusOK)
		} else {
			c.handleRequest(w, r)
			next.ServeHTTP(w, r)
		}
	})
}

func (c CORS) handleOptions(w http.ResponseWriter, r *http.Request) {
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

func (c CORS) handleRequest(w http.ResponseWriter, r *http.Request) {
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

func (c CORS) isOriginAllowed(origin string) bool {
	if !c.checkOrigin {
		return true
	}

	if _, ok := c.originsList[strings.ToLower(origin)]; ok {
		return true
	}

	return false
}

func (c CORS) isMethodAllowed(method string) bool {
	if !c.checkMethods {
		return true
	}

	if _, ok := c.methodsList[strings.ToUpper(method)]; ok {
		return true
	}

	return false
}
