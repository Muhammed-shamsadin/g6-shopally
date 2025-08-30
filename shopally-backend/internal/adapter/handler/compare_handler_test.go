package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompareProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success Case: Should return 200 OK with comparison data", func(t *testing.T) {
		// Arrange
		mockUseCase := new(usecase.MockCompareProductsUseCase)
		handler := NewCompareHandler(mockUseCase)

		router := gin.Default()
		router.POST("/compare", handler.CompareProducts)

		productsToCompare := []*domain.Product{
			{ID: "ALI-123", Title: "Product A"},
			{ID: "ALI-456", Title: "Product B"},
		}
		requestBody, _ := json.Marshal(gin.H{"products": productsToCompare})

		expectedResult := map[string]interface{}{"comparison": "some comparison data"}

		mockUseCase.
			On("Execute", mock.Anything, productsToCompare).
			Return(expectedResult, nil).
			Once()

		// Act
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/compare", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &responseBody)

		// Use EqualValues to ignore type alias differences
		assert.EqualValues(t, expectedResult, responseBody["data"])
		assert.Nil(t, responseBody["error"])

		mockUseCase.AssertExpectations(t)
	})

	t.Run("Validation Error Case: Should return 400 Bad Request for invalid number of products", func(t *testing.T) {
		// Arrange
		mockUseCase := new(usecase.MockCompareProductsUseCase)
		handler := NewCompareHandler(mockUseCase)

		router := gin.Default()
		router.POST("/compare", handler.CompareProducts)

		productsToCompare := []*domain.Product{
			{ID: "ALI-123", Title: "Product A"},
		}
		requestBody, _ := json.Marshal(gin.H{"products": productsToCompare})

		// Act
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/compare", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &responseBody)

		errorData := responseBody["error"].(map[string]interface{})
		assert.Equal(t, "INVALID_INPUT", errorData["code"])
		assert.Contains(t, errorData["message"], "must contain a 'products' array with 2 to 4")

		// The use case should never be called
		mockUseCase.AssertNotCalled(t, "Execute")
	})

	t.Run("Malformed JSON Case: Should return 400 Bad Request for bad JSON", func(t *testing.T) {
		// Arrange
		mockUseCase := new(usecase.MockCompareProductsUseCase)
		handler := NewCompareHandler(mockUseCase)

		router := gin.Default()
		router.POST("/compare", handler.CompareProducts)

		invalidRequestBody := []byte(`{"products": [,]}`)

		// Act
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/compare", bytes.NewBuffer(invalidRequestBody))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// The use case should never be called
		mockUseCase.AssertNotCalled(t, "Execute")
	})
}
