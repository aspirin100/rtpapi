package service

import (
	"math/rand"
	"time"

	"github.com/aspirin100/rtpapi/internal/config"
)

const (
	MinValue float32 = 1
	MaxValue float32 = 10000
)

type Service struct {
	Rtp       *float32
	generator *rand.Rand
}

func New(cfg *config.Config) *Service {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	return &Service{
		Rtp:       &cfg.Rtp,
		generator: r,
	}
}

func (s *Service) GenerateMultiplier() float32 {
	u := s.generator.Float32()

	// generate high mutipliers
	if *s.Rtp == 1.0 {
		return MinValue + s.generator.Float32()*(MaxValue-MinValue)
	}

	threshold := MinValue + (*s.Rtp)*(MaxValue-MinValue)

	if u <= *s.Rtp {
		return threshold + s.generator.Float32()*(MaxValue-threshold)
	}

	// generate small multiplier
	return MinValue + s.generator.Float32()*(threshold-MinValue)
}

// for test
func (s *Service) ChangeSeed() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	s.generator = r
}
