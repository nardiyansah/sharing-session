package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-key") != "secret" {
			w.Write([]byte("request denied\n"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world\n")
}

func greeting(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body := make(map[string]string)
		d := json.NewDecoder(r.Body)

		err := d.Decode(&body)
		if err != nil {
			fmt.Fprint(w, "error")
		}

		fmt.Fprintf(w, "hi %v\n", body["name"])
	}
}

func main() {

	mux := http.DefaultServeMux

	mux.HandleFunc("/", hello)
	mux.HandleFunc("/greeting", greeting)

	var handler http.Handler = mux
	handler = middleware(handler)

	server := new(http.Server)
	server.Addr = ":8000"
	server.Handler = handler

	fmt.Printf("server listen and serve at %v\n", server.Addr)
	server.ListenAndServe()
}
