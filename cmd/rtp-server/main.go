package main

import (
	"fmt"
	"log"

	"github.com/aspirin100/rtpapi/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
