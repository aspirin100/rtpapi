package service

import (
	"math/rand"
)

const (
	MinValue float32 = 1
	MaxValue float32 = 10000
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GenerateMultiplier() float32 {
	return MinValue + rand.Float32()*(MaxValue-MinValue)
}
