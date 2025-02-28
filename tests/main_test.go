package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSignerRecord(t *testing.T) {
	t.Parallel()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method, "Expected POST request")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"record_id": "12345"}`))
	})

	reqBody := bytes.NewBuffer([]byte(`{"signer": "Teszt Ellek"}`))
	req, err := http.NewRequest(http.MethodPost, "/create", reqBody)
	require.NoError(t, err, "Failed to create request")

	respRecorder := httptest.NewRecorder()
	handler.ServeHTTP(respRecorder, req)

	assert.Equal(t, http.StatusOK, respRecorder.Code, "Expected 200 OK response")

	var response struct {
		RecordID string `json:"record_id"`
	}
	err = json.NewDecoder(respRecorder.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, "12345", response.RecordID, "Expected record_id to be 12345")
}

func TestUploadPDF(t *testing.T) {
	t.Parallel()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method, "Expected POST request")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "File uploaded successfully"})
	})

	reqBody := bytes.NewBuffer([]byte("dummy pdf content"))
	req, err := http.NewRequest(http.MethodPost, "/upload", reqBody)
	require.NoError(t, err, "Failed to create request")

	respRecorder := httptest.NewRecorder()
	handler.ServeHTTP(respRecorder, req)

	assert.Equal(t, http.StatusOK, respRecorder.Code, "Expected 200 OK response")

	var response struct {
		Message string `json:"message"`
	}
	err = json.NewDecoder(respRecorder.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, "File uploaded successfully", response.Message, "Unexpected response message")
}

func TestFinalizeRecord(t *testing.T) {
	t.Parallel()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method, "Expected POST request")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "finalized"}`))
	})

	reqBody := bytes.NewBuffer([]byte(`{"record_id": "12345"}`))
	req, err := http.NewRequest(http.MethodPost, "/finalize", reqBody)
	require.NoError(t, err, "Failed to create request")

	respRecorder := httptest.NewRecorder()
	handler.ServeHTTP(respRecorder, req)

	assert.Equal(t, http.StatusOK, respRecorder.Code, "Expected 200 OK response")

	var response struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(respRecorder.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, "finalized", response.Status, "Expected status to be 'finalized'")
}
