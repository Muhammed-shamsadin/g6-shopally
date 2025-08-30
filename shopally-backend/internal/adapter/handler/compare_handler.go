package handler

import (
	"net/http"

	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/usecase"

	"github.com/gin-gonic/gin"
)

// CompareHandler is responsible for handling all HTTP requests related to product comparison.
type CompareHandler struct {
	compareUseCase *usecase.CompareProductsUseCase
}

// NewCompareHandler creates a new instance of the CompareHandler.
func NewCompareHandler(uc *usecase.CompareProductsUseCase) *CompareHandler {
	return &CompareHandler{
		compareUseCase: uc,
	}
}

// CompareProducts is the Gin handler function for the POST /compare endpoint.
func (h *CompareHandler) CompareProducts(c *gin.Context) {

	var requestBody struct {
		Products []*domain.Product `json:"products"`
	}

	// Try to bind the incoming JSON to our struct.
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		// If binding fails, it's a client error (malformed JSON).
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INVALID_INPUT",
				"message": "Invalid request body. Ensure it is a valid JSON.",
			},
		})
		return
	}

	// check if the products are between two and four inclusively
	if len(requestBody.Products) < 2 || len(requestBody.Products) > 4 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INVALID_INPUT",
				"message": "Request body must contain a 'products' array with 2 to 4 product objects.",
			},
		})
		return
	}

	comparisonResult, err := h.compareUseCase.Execute(c.Request.Context(), requestBody.Products)
	if err != nil {
		// If the use case returns an error, it's a server-side issue.
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "An error occurred while comparing products.",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  comparisonResult,
		"error": nil,
	})
}
