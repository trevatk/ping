package main

import (
	"log"
	"net/http"
	"os"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Printf("received health check from %s \n", r.Host)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("failed to write response %v", err)
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", health)

	port := os.Getenv("SERVER_HTTP_PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("start http server http://127.0.0.1:%s", port)
	http.ListenAndServe(":"+port, nil)
}
