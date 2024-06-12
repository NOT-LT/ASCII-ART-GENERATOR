package asciiweb

import (
	asciiweb "asciiweb/internal/ascii-art"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		asciiweb.ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		asciiweb.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		asciiweb.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		asciiweb.ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	input := r.FormValue("input")
	if len(input) > 975 {
		asciiweb.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")
	AsciiResult, err := asciiweb.AsciiArt(input, banner)
	if err != nil {
		// AsciiResult, _ = AsciiWeb.AsciiArt("Sorry Invalid :(", banner)
		asciiweb.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	if AsciiResult == "" {
		asciiweb.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	t, _ := template.ParseFiles("static/index.html")
	if AsciiResult != "" {
		t.Execute(w, AsciiResult)
	}
}

// Read about http headings guys :>
func Downloadhandler(w http.ResponseWriter, r *http.Request, input string) {
	fmt.Println(len(input))
	// if len(input) > 500 {
	// 	// handle error HERE
	// 	return
	// }
	downloadButton := r.FormValue("downloadButton")

	if downloadButton == "" {
		// handle the error here
		// مادري وش نوع الايرور
		return
	}

	// Content Disposition -> How to display the content
	// here it's an attachment (Downloaded)
	w.Header().Set("Content-Disposition", "attachment; filename=yay.txt")

	// indicate the media type, here it's text
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(downloadButton)))
	// Write the AsciiResult to the response
	w.Write([]byte(downloadButton))
}
