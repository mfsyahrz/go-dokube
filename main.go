package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is empty")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "bad request", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("hello world"))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":" + port

	log.Println("server up at: ", port)
	log.Println(server.ListenAndServe())
}
