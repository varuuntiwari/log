package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	port := os.Getenv("SERVER_PORT")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		ip := r.Header.Get("X-Real-IP")
		fmt.Fprintf(w, "<html><head><title>Your IP</title></head><body>%v</body></html>", ip)
	}).Methods("GET")
	
	http.Handle("/", router)
	http.ListenAndServe(":" + port, router)
}