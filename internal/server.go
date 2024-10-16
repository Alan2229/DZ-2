package internal

import (
	"internal/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/version", handlers.Version)
	mux.HandleFunc("/decode", handlers.Decode)
	mux.HandleFunc("/hard-op", handlers.HardOp)
	server := &http.Server{Addr: ":8080", Handler: mux}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("", err)
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("EREOROEO")
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
