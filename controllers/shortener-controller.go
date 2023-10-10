package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type TemplatePayload struct {
	Result   string
	Err      string
	Host     string
	Protocol string
}

func generateHash(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var shortLink entities.ShortLink

	shortLink.Origin = r.FormValue("origin")
	shortLink.Short = r.FormValue("short")

	if shortLink.Short == "" {
		shortLink.Short = generateHash(7)
	}

	tml := template.Must(template.ParseFiles("./static/generate.html"))

	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}

	result := database.Instance.Create(&shortLink)
	if result.Error != nil {
		fmt.Println(result.Error)
		tml.Execute(w, TemplatePayload{"", result.Error.Error(), r.Host, protocol})
		return
	}

	err := tml.Execute(w, TemplatePayload{shortLink.Short, "", r.Host, protocol})
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}

func GetOrigin(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	var shortLink entities.ShortLink
	if err := database.Instance.Where("Short = ?", hash).First(&shortLink).Error; err != nil {
		fmt.Println(err)
		respondJSON(w, http.StatusNotFound, ErrorResponse{fmt.Sprintf("Link Not found with ID: %s", hash)})
		return
	}

	http.Redirect(w, r, shortLink.Origin, http.StatusSeeOther)
}
