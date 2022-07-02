package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world\n")
	})

	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body := make(map[string]string)
			d := json.NewDecoder(r.Body)

			err := d.Decode(&body)
			if err != nil {
				fmt.Fprint(w, "error")
			}

			fmt.Fprintf(w, "hi %v\n", body["name"])
		}
	})

	http.ListenAndServe(":8000", nil)
}
