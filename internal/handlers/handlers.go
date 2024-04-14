package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/DiegusNueva/DiegoAlonso_Website_Go/internal/models"
)

func renderTemplate(w http.ResponseWriter, tmplFile string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title:   "TuniverS",
		Author:  "Diego Alonso Molina",
		Welcome: "Mid Full Stack Developer | React | MySQL | Go | HTML5 | CSS3 | JavaScript e Ingeniero Técnico en Topografía (Surveying Engineer)",
	}

	page := r.URL.Path[1:]

	if page == "" {
		page = "index.html"
	}

	tmplFile := "web/templates/" + page

	// Head functions
	headContent, err := os.ReadFile("web/templates/head.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	headTemplate, err := template.New("head").Parse(string(headContent))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var headBuffer bytes.Buffer
	err = headTemplate.Execute(&headBuffer, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data.HeadContent = template.HTML(headBuffer.String())

	// Footer functions
	footerContent, err := os.ReadFile("web/templates/footer.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	footerTemplate, err := template.New("footer").Parse(string(footerContent))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var footerBuffer bytes.Buffer
	err = footerTemplate.Execute(&footerBuffer, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data.FooterContent = template.HTML(footerBuffer.String())

	// Navbar functions
	navbarContent, err := os.ReadFile("web/templates/navbar.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	navbarTemplate, err := template.New("navbar").Parse(string(navbarContent))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var navbarBuffer bytes.Buffer
	err = navbarTemplate.Execute(&navbarBuffer, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data.NavbarContent = template.HTML(navbarBuffer.String())

	if _, err := os.Stat(tmplFile); err != nil {
		tmplFile = "web/templates/error.html"

		data.ErrorCode = http.StatusNotFound
		data.ErrorMessage = "Página no encontrada"
	}

	renderTemplate(w, tmplFile, data)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title:        "¡Página no encontrada",
		ErrorCode:    http.StatusInternalServerError,
		ErrorMessage: "Error interno del servidor",
	}

	renderTemplate(w, "web/templates/error.html", data)
}