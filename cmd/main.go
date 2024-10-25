package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello")

	// TODO: use some router lib instead
	http.HandleFunc("/bid", handleBidder)
	http.HandleFunc("/win", handleWin)
}

func handleBidder(w http.ResponseWriter, r *http.Request) {
	// TODO: handle

}

func handleWin(w http.ResponseWriter, r *http.Request) {
	// TODO: implement

}
