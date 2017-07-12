package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Print("GET / 200 OK")
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":5000", nil)
}
