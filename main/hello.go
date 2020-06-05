package main

import (
	"fmt"
	"github.com/R2wenD2/cb"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, fmt.Sprintf("%d", cb.Add(1,1)))
        })


	log.Fatal(http.ListenAndServe(":8081", nil))

}
