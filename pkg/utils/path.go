package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DefinePath(router *mux.Router) {
	staticFileDirectory := http.Dir("web/static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")
}
