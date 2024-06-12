package main

import (
	asciiweb "asciiweb/internal/Handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", asciiweb.GetHandler)
	http.HandleFunc("/ascii-art", asciiweb.PostHandler)
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		// Additional parameters for the Downloadhandler function
		input := r.FormValue("input")
		fmt.Println(input)
		param1 := input

		asciiweb.Downloadhandler(w, r, param1)
	})
	// Without it css won't be applied
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/styles.css")
	})
	http.HandleFunc("/error.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/error.css")
	})
	port := "localhost:8080"
	fmt.Printf("Server is working on http://" + port + "\n")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {

		}
	}()
	// nil means DefaultServeMux is used
	asciiweb.OpenBrowser("http://localhost:8080")
	select {}
}
