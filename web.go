package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/tokens", tokensHandler)
	r.HandleFunc("/token/{symbol}", tokenHandler)
	r.Use(apiMiddleware)
	log.Fatal(http.ListenAndServe(serverHost, r))
}

func apiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

// Returns all tokens
func tokensHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(tokens); err != nil {
		log.Error(err)
		http.Error(w, "Server Error", 500)
		return
	}
}

// Returns info for specific token by symbol
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := strings.ToUpper(vars["symbol"])

	token := tokensBySymbol[symbol]
	if token == nil {
		http.Error(w, "Not Found", 404)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(token); err != nil {
		log.Error(err)
		http.Error(w, "Server Error", 500)
		return
	}
}
