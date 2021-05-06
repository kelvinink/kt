package main

import (
	"log"
	"net/http"
	"time"
)

type Handler func(http.ResponseWriter, *http.Request)

type Middleware func(Handler) Handler

// Chain multi middleware to one
func Chain(mws ...Middleware) Middleware {
	return func(handler Handler) Handler {
		for _, mw := range mws {
			handler = mw(handler)
		}
		return handler
	}
}

// Logging middleware
func Logging(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[access] %s", r.URL.String())
		handler(w, r)
	}
}

func Date(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[current time] %s", time.Now().String())
		handler(w, r)
	}
}

// Hello handler
func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello man"))
}

// Hi handler
func Hi(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hi man"))
}

func main() {
	http.HandleFunc("/hello", Logging(Hello))
	http.HandleFunc("/hi", Logging(Hi))
	http.HandleFunc("/chain", Chain(Logging, Date)(Hi))

	log.Panic(http.ListenAndServe(":8080", nil))
}
