package bidder

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"pet-bidder/internal/data"
	"pet-bidder/internal/storage"
)

type Server interface {
	Run() error
	Close() error
	FillStorage() error
}

type PetBidderService struct {
	storage    storage.Storage
	httpServer *http.Server
}

// NewServer creates a new Server using given protocol and addr.
func NewServer(addr string) (Server, error) {
	localStore := storage.NewLocalStorage()
	srv := &PetBidderService{
		storage: localStore,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/bid", srv.handleBidRequest)
	mux.HandleFunc("/win", srv.handleWin)

	srv.httpServer = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return srv, nil
}

func (s *PetBidderService) Run() error {
	log.Printf("Lisnet and Serve on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *PetBidderService) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Shutting down server gracefully...")
	return s.httpServer.Shutdown(ctx)
}

func (s *PetBidderService) FillStorage() error {
	campaigns, err := data.GetInitialData()
	if err != nil {
		return fmt.Errorf("failed on GetInitialData: %v\n", err)
	}

	log.Printf("Fake campaigns: %+v\n", campaigns)

	err = s.storage.BulkSet(campaigns)
	return err
}
