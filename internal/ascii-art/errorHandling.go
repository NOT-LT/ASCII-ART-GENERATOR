package asciiweb

import (
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status) // to set the status code
	parsedTemplate, err := template.ParseFiles("static/404.html")
	if status == http.StatusNotFound {
		parsedTemplate, err = template.ParseFiles("static/404.html")
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	} else if status == http.StatusBadRequest {
		parsedTemplate, err = template.ParseFiles("static/400.html")
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	} else if status == http.StatusInternalServerError {
		parsedTemplate, _ = template.ParseFiles("static/500.html")
	} else if status == http.StatusMethodNotAllowed {
		parsedTemplate, err = template.ParseFiles("static/405.html")
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}
	parsedTemplate.Execute(w, nil)
}
