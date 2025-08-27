package gateway

import (
	"context"

	"github.com/shopally-ai/pkg/usecase"
)

// MockLLMGateway implements usecase.LLMGateway and returns a hardcoded parsed intent.
type MockLLMGateway struct{}

func NewMockLLMGateway() usecase.LLMGateway {
	return &MockLLMGateway{}
}

func (m *MockLLMGateway) ParseIntent(ctx context.Context, query string) (map[string]interface{}, error) {
	// Very simple mocked intent
	return map[string]interface{}{
		"category":      "smartphone",
		"price_max_ETB": 5000,
	}, nil
}
