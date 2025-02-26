package service_test

import (
	"fmt"
	"testing"

	"github.com/aspirin100/rtpapi/internal/config"
	"github.com/aspirin100/rtpapi/internal/service"
)

const (
	commonTestCount         = 20000
	MinSequanceVal  float32 = 1
	MaxSequenceVal  float32 = 10000
)

func TestGenerateMultiplier(t *testing.T) {
	// Service setup
	var randRtp float32 = 1

	cfg := config.Config{
		Rtp: float32(randRtp),
	}

	srv := service.New(&cfg)
	srv.ChangeSeed()

	// Client setup
	sq := []float32{2.0, 3.0, 5.0, 1.1, 1.1}
	sum0 := len(sq)

	var totalSum1 float32 = 0

	temp := make([]float32, sum0)

	for i := 0; i < commonTestCount; i++ {
		if i == 0 {
			copy(temp, sq)
		}

		var sum1 float32 = 0
		for j := 0; j < sum0; j++ {
			mp := srv.GenerateMultiplier()
			if mp <= temp[j] {
				temp[j] = 0
			}
			sum1 += temp[j]
		}
		totalSum1 += sum1
	}

	averageRTP := totalSum1 / (float32(commonTestCount) * float32(sum0))
	fmt.Printf("Average RTP: %f\n", averageRTP)

	if averageRTP < randRtp-0.05 || averageRTP > randRtp+0.05 {
		t.Errorf("RTP is not close to the target value. Expected: %f, Got: %f", randRtp, averageRTP)
	}
}
