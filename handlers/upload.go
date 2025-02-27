package handlers

import (
	"Comnica_SignIN_task/services"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

// RenderForm megjeleníti az aláírási formot
func RenderForm(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

// UploadFile kezeli a PDF feltöltést és az aláírási folyamatot
func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Fájl beolvasása a kérésből
	file, _, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to retrieve file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ideiglenes fájl létrehozása
	tempFile, err := os.CreateTemp("./data", "upload-*.pdf")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create temporary file: %v", err), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, file); err != nil {
		http.Error(w, fmt.Sprintf("failed to save uploaded file: %v", err), http.StatusInternalServerError)
		return
	}

	// Session létrehozása és dokumentum feltöltése
	sessionID, token, err := services.CreateSession()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create session: %v", err), http.StatusInternalServerError)
		return
	}

	signURL, documentID, err := services.UploadDocument(sessionID, token, tempFile.Name())
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to upload document: %v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("Session ID: %s, Document ID: %d", sessionID, documentID)

	// Visszajelzés a kliens felé
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}
	data := map[string]string{
		"SignURL":   signURL,
		"SessionID": sessionID,
	}
	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
	}
}
