package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"gluster/utils"
)

// RestLoggingHandler is a Middleware to log the incoming request as
// per the Apache Common logging framework
func RestLoggingHandler(h http.Handler) http.Handler {
	logFile, err := os.OpenFile(utils.RestConfig.AccessLogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, h)
}

// SetApplicationHeaderJSON is a Middleware to set Application
// Header as JSON
func SetApplicationHeaderJSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}
