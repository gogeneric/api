package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gogeneric/api"
)

func main() {
	h := &http.Server{Addr: "0.0.0.0:3000"}
	mux := http.NewServeMux()
	h.Handler = mux

	// Here is magic!
	mux.Handle("/hello", api.Wrap(handleHello))

	if err := h.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func handleHello(ctx context.Context, req *helloRequest) (*helloResponse, error) {
	return &helloResponse{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string `json:"message"`
}
