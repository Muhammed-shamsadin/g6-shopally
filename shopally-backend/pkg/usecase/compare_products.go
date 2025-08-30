package usecase

import (
	"context"

	"github.com/shopally-ai/pkg/domain"
)

// compare products
type CompareProductsUseCase struct {
	llmGateway domain.LLMGateway
}

func (uc *CompareProductsUseCase) Execute(context context.Context, products []*domain.Product) (any, any) {
	panic("unimplemented")
}

func NewCompareProductsUseCase(lg domain.LLMGateway) *CompareProductsUseCase {
	return &CompareProductsUseCase{
		llmGateway: lg,
	}
}

func (uc *CompareProductsUseCase) Compare(ctx context.Context, query string) (interface{}, error) {
	// Parse intent via LLM
	_, err := uc.llmGateway.CompareProducts(ctx, []*domain.Product{})
	if err != nil {
		// For V1 mock, fail soft by using empty filters
		_ = map[string]interface{}{}
	}
	// Return the envelope-compatible data payload
	return map[string]interface{}{"products": ""}, nil
}
