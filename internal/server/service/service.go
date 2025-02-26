package service

import (
	"math/rand"

	"github.com/aspirin100/rtpapi/internal/config"
)

type Service struct {
	MinValue *float32
	MaxValue *float32
}

func New(cfg *config.Config) *Service {
	return &Service{
		MinValue: &cfg.MinValue,
		MaxValue: &cfg.MaxValue,
	}
}

func (s *Service) GenerateMultiplier() float32 {
	return *s.MinValue + rand.Float32()*(*s.MaxValue-*s.MinValue)
}
