package main

import (
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

type Middleware func(Handler) Handler

// Router config withe middleware and handler
type Router struct {
	chain []Middleware
	mux   map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		chain: []Middleware{},
		mux:   make(map[string]Handler),
	}
}

// Use a middleware
func (r *Router) Use(m Middleware) {
	r.chain = append(r.chain, m)
}

// AddRoute config
func (r *Router) AddRoute(route string, h Handler) {
	var mergedHandler = h

	for _, mw := range r.chain {
		mergedHandler = mw(mergedHandler)
	}

	r.mux[route] = mergedHandler
}

// Activate Route
func (r *Router) ActivateRoute() {
	for k, v := range r.mux {
		http.HandleFunc(k, v)
	}
}

// Chain multi middleware to one
func Chain(mws ...Middleware) Middleware {
	return func(h Handler) Handler {
		for _, mw := range mws {
			h = mw(h)
		}
		return h
	}
}

// Logger middleware
func Logger(h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[logger]: %s", r.URL.String())
		h(w, r)
	}
}

// Profiler middleware
func Profiler(h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[profiler]: %s", "mem: 50%, cpu: 50%")
		h(w, r)
	}
}

// Hello handler
func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("say hello..."))
}

func main() {
	http.HandleFunc("/hello", Logger(Hello))
	http.HandleFunc("/chain", Chain(Logger, Profiler)(Hello))

	r := NewRouter()
	r.Use(Logger)
	r.Use(Profiler)
	r.AddRoute("/use", Hello)
	r.ActivateRoute()

	log.Panic(http.ListenAndServe(":8080", nil))
}
