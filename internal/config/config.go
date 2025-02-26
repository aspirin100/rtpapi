package config

import (
	"errors"
	"flag"
	"fmt"
)

var (
	ErrNoRtp = errors.New("rtp flag must be set")
)

type Config struct {
	Rtp float32
}

func New() (*Config, error) {
	rtp, err := fetchMultiplierFlag()
	if err != nil {
		return nil, fmt.Errorf("configuation fail: %w", err)
	}

	return &Config{
		Rtp: *rtp,
	}, nil
}

func fetchMultiplierFlag() (*float32, error) {
	var rtp float64

	flag.Float64Var(&rtp, "rtp", 0, "")
	flag.Parse()

	if rtp == 0 {
		return nil, ErrNoRtp
	}

	rtp32 := float32(rtp)

	return &rtp32, nil
}
