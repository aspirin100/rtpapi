package service

import (
	"math/rand"

	"github.com/aspirin100/rtpapi/internal/config"
)

const (
	MinValue float32 = 1
	MaxValue float32 = 10000
)

type Service struct {
	Rtp *float32
}

func New(cfg *config.Config) *Service {
	return &Service{
		Rtp: &cfg.Rtp,
	}
}

func (s *Service) GenerateMultiplier() float32 {
	const lowMultiplier = 10

	u := rand.Float64()

	randVal := float32(rand.Float64())

	if u <= float64(*s.Rtp) {
		return MinValue + randVal*(MaxValue-MinValue)
	}

	return MinValue + randVal*lowMultiplier
}
