package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"pet-bidder/internal/bidder"
	"pet-bidder/internal/config"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("Failed on parse config: %v\n", err)
	}

	serv, err := bidder.NewServer(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("Failed on init server: %v\n", err)
		return
	}

	err = serv.FillStorage()
	if err != nil {
		log.Fatalf("Failed on fill storage with initial data: %v\n", err)
		return
	}

	// Listen to Unix signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-sig
		log.Printf("Received signal %s. Close the Server.\n", s.String())
		err := serv.Close()
		if err != nil {
			log.Printf("Server failed on Close: %v\n", err)
		}
	}()

	log.Println("Server Running...")
	err = serv.Run()
	if err != nil {
		log.Fatalf("Server Error: %v\n", err)
	}
}
