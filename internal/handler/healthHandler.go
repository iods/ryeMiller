package handler

import (
	"html/template"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dir, _ := os.Getwd()
	templates := template.Must(template.ParseFiles(dir + "/web/template/page/health.gohtml"))

	if err := templates.ExecuteTemplate(w, "health.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}