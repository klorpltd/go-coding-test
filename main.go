package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/klorpltd/go-coding-test/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)
		<-sigCh
		cancel()
	}()

	if err := server.Run(ctx, 8080); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server.Run failed: %v\n", err)
	}
}
