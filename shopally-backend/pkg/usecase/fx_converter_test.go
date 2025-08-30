package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/shopally-ai/internal/mocks"
	"github.com/stretchr/testify/suite"
)

type PriceConverterSuite struct {
	suite.Suite
	ctx context.Context
	fx  *mocks.IFXClient
	pc  PriceConverter
}

func (s *PriceConverterSuite) SetupTest() {
	s.ctx = context.Background()
	s.fx = mocks.NewIFXClient(s.T())
	s.pc = PriceConverter{FX: s.fx}
}

func (s *PriceConverterSuite) TestUSDToETB_Success() {
	s.fx.On("GetRate", s.ctx, "USD", "ETB").Return(56.5, nil).Once()
	etb, rate, err := s.pc.USDToETB(s.ctx, 10.0)
	s.Require().NoError(err)
	s.InDelta(56.5, rate, 1e-9)
	s.InDelta(565.0, etb, 1e-9)
}

func (s *PriceConverterSuite) TestUSDToETB_Error() {
	s.fx.On("GetRate", s.ctx, "USD", "ETB").Return(0.0, errors.New("fx down")).Once()
	etb, rate, err := s.pc.USDToETB(s.ctx, 5.0)
	s.Error(err)
	s.Equal(0.0, rate)
	s.Equal(0.0, etb)
}

func TestPriceConverterSuite(t *testing.T) { suite.Run(t, new(PriceConverterSuite)) }
