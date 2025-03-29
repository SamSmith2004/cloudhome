package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func HelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure /temp directory exists
		if _, err := os.Stat("./temp"); os.IsNotExist(err) {
			os.Mkdir("./temp", 0755)
		}

		// Ensure content is a  file and <= 10MB
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		tempPath := filepath.Join("./temp", handler.Filename)
		dst, err := os.Create(tempPath)
		if err != nil {
			http.Error(w, "Error creating file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("success"))
	}
}
