package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"Comnica_SignIN_task/config"
)

// AddDocumentResponse a dokumentumfeltöltés válaszstruktúrája
type AddDocumentResponse struct {
	DocumentID int `json:"document_id"`
}

// UploadDocument feltölti a PDF-et az aláírási sessionhöz
func UploadDocument(sessionID, token, filename string) (string, int, error) {
	// Fájl megnyitása
	file, err := os.Open(filename)
	if err != nil {
		return "", 0, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Multipart request összeállítása
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("data", filepath.Base(filename))
	if err != nil {
		return "", 0, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", 0, fmt.Errorf("failed to copy file content: %w", err)
	}
	if err := writer.WriteField("session_id", sessionID); err != nil {
		return "", 0, fmt.Errorf("failed to write session_id field: %w", err)
	}
	if err := writer.WriteField("description", "User Document"); err != nil {
		return "", 0, fmt.Errorf("failed to write description field: %w", err)
	}
	if err := writer.WriteField("filename", filepath.Base(filename)); err != nil {
		return "", 0, fmt.Errorf("failed to write filename field: %w", err)
	}
	if err := writer.Close(); err != nil {
		return "", 0, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// HTTP kérelem összeállítása és elküldése
	url := fmt.Sprintf("%s/session/add_document", config.TestURL)
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("document upload failed: %s", resp.Status)
	}

	// Válasz dekódolása
	var res AddDocumentResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", 0, fmt.Errorf("failed to decode response: %w", err)
	}

	signURL := fmt.Sprintf("https://sign-test.comnica.com/sign/%s", sessionID)
	return signURL, res.DocumentID, nil
}
