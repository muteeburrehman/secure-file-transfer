package handlers

import (
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	filePath := "./static/uploads/" + filename
	http.ServeFile(w, r, filePath)
}
