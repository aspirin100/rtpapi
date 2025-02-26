package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aspirin100/rtpapi/internal/config"
	"github.com/aspirin100/rtpapi/internal/server/handler"
	"github.com/aspirin100/rtpapi/internal/server/service"
)

const (
	shutdownTimeout = time.Second * 5
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("current config", cfg)

	srv := service.New()

	application := handler.New(srv)

	go func() {
		err = application.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	stopCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	err = application.Shutdown(stopCtx)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("correctly stopped")
}
