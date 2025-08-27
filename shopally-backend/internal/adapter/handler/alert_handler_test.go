package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopally-ai/internal/adapter/repository"
	"github.com/shopally-ai/pkg/usecase"
)

func TestAlertHandlers(t *testing.T) {
	mockRepo := repository.NewMockAlertRepository()
	alertManager := usecase.NewAlertManager(mockRepo)
	alertHandler := NewAlertHandler(alertManager)

	var alertID string

	t.Run("CreateAlertHandler", func(t *testing.T) {
		payload := []byte(`{"userId": "user-123", "productId": "prod-abc", "targetPrice": 500.00}`)

		req := httptest.NewRequest("POST", "/alerts", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		alertHandler.CreateAlertHandler(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		var res successResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}

		dataMap, ok := res.Data.(map[string]interface{})
		if !ok {
			t.Fatalf("response data is not a map: got %T", res.Data)
		}
		if status, ok := dataMap["status"].(string); !ok || status != "Alert created successfully" {
			t.Errorf("unexpected status message: got %v", status)
		}

		dataMap, ok = res.Data.(map[string]interface{})
		if !ok {
			t.Errorf("response data is not a map: got %T", res.Data)
		} else {
			id, ok := dataMap["alertId"].(string)
			if !ok || id == "" {
				t.Errorf("missing or invalid alertId in response: got %v", id)
			} else {
				alertID = id
			}
		}
	})

	t.Run("GetAlertHandler", func(t *testing.T) {
		if alertID == "" {
			t.Fatal("alertID was not set in previous test")
		}

		req := httptest.NewRequest("GET", "/alerts/"+alertID, nil)
		rr := httptest.NewRecorder()

		alertHandler.GetAlertHandler(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var res successResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}

		if alertData, ok := res.Data.(map[string]interface{}); !ok || alertData["alertId"] != alertID {
			t.Errorf("unexpected alertId in response: got %v", alertData["alertId"])
		}
	})

	t.Run("DeleteAlertHandler", func(t *testing.T) {
		if alertID == "" {
			t.Fatal("alertID was not set in previous test")
		}

		req := httptest.NewRequest("DELETE", "/alerts/"+alertID, nil)
		rr := httptest.NewRecorder()

		alertHandler.DeleteAlertHandler(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var res successResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}

		dataMap, ok := res.Data.(map[string]interface{})
		if !ok {
			t.Fatalf("response data is not a map: got %T", res.Data)
		}
		if status, ok := dataMap["status"].(string); !ok || status != "Alert deleted successfully" {
			t.Errorf("unexpected status message: got %v", status)
		}
	})

	t.Run("GetDeletedAlertFails", func(t *testing.T) {
		if alertID == "" {
			t.Fatal("alertID was not set in previous test")
		}

		req := httptest.NewRequest("GET", "/alerts/"+alertID, nil)
		rr := httptest.NewRecorder()

		alertHandler.GetAlertHandler(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}
