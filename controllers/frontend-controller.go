package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupIndex(router *mux.Router) {
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/").Handler(fs)
}