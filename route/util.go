package route

import (
	"net/http"
	"fmt"
	"html/template"
)

const(
	layoutFN = "layout"
)

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string){
	var files []string
	for _, file := range fn{
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, layoutFN, data)
}
