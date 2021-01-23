package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mainCtx, cancel := context.WithCancel(context.Background())

	go watchSignals(cancel)

	server := newServer()

	go func() {
		log.Println("Starting server...")
		err := server.Start(":8000")
		if err != nil {
			log.Print(err)
			cancel()
		}
	}()

	<-mainCtx.Done()

	cancel()

	log.Println("Stopping server")
	shutDown(server)
	log.Println("Server is stopped")
}

func watchSignals(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	cancel()
}

func shutDown(server *Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		log.Print(err)
	}
}
