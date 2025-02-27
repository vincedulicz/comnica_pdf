package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"Comnica_SignIN_task/config"
)

// InitSessionResponse a session létrehozásának válaszstruktúrája
type InitSessionResponse struct {
	BearerToken string `json:"bearer_token"`
	SessionID   string `json:"session_id"`
}

// CreateSession létrehozza az aláírási sessiont
func CreateSession() (string, string, error) {
	payload := map[string]string{
		"company": config.Company,
		"case_id": "example-case-123",
		"name":    "Test User",
		"email":   "test@example.com",
		"phone":   "36201234567",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := fmt.Sprintf("%s/session/init", config.TestURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+config.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("session initialization failed: %s", resp.Status)
	}

	var res InitSessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", "", fmt.Errorf("failed to decode response: %w", err)
	}
	return res.SessionID, res.BearerToken, nil
}
