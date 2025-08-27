// Package handler provides HTTP handlers for alert-related endpoints.
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/usecase"
)

// AlertHandler handles HTTP requests for alert operations.
type AlertHandler struct {
	alertManager *usecase.AlertManager
}

// NewAlertHandler creates a new AlertHandler with the given AlertManager.
func NewAlertHandler(am *usecase.AlertManager) *AlertHandler {
	return &AlertHandler{
		alertManager: am,
	}
}

// successResponse represents a standard API response structure.
type successResponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

// createAlertPayload represents the expected payload for creating an alert.
type createAlertPayload struct {
	UserID      string  `json:"userId"`
	ProductID   string  `json:"productId"`
	TargetPrice float64 `json:"targetPrice"`
}

// CreateAlertHandler handles POST requests to create a new alert.
func (h *AlertHandler) CreateAlertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload createAlertPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newAlert := &domain.Alert{
		UserID:      payload.UserID,
		ProductID:   payload.ProductID,
		TargetPrice: payload.TargetPrice,
		IsActive:    true,
	}

	if err := h.alertManager.CreateAlert(newAlert); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create alert: %v", err), http.StatusInternalServerError)
		return
	}

	response := successResponse{
		Data: map[string]string{
			"status":  "Alert created successfully",
			"alertId": newAlert.ID,
		},
		Error: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GetAlertHandler handles GET requests to retrieve an alert by its ID.
func (h *AlertHandler) GetAlertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}
	alertID := parts[2]

	alert, err := h.alertManager.GetAlert(alertID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve alert: %v", err), http.StatusNotFound)
		return
	}

	response := successResponse{
		Data:  alert,
		Error: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// DeleteAlertHandler handles DELETE requests to remove an alert by its ID.
func (h *AlertHandler) DeleteAlertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}
	alertID := parts[2]

	if err := h.alertManager.DeleteAlert(alertID); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete alert: %v", err), http.StatusNotFound)
		return
	}

	response := successResponse{
		Data: map[string]string{
			"status": "Alert deleted successfully",
		},
		Error: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
