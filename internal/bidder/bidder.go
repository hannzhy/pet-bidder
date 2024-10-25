package bidder

import (
	"context"
	"log"
	"net/http"
	"pet-bidder/internal/storage"
	"time"
)

type Server interface {
	Run() error
	Close() error
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
