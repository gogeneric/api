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
	User string
}

// SetHeaders implements api.WithHeader interface
func (h *helloRequest) WithHeader(header http.Header) {
	h.User = header.Get("user")
}

type helloResponse struct {
	Message string `json:"message"`
}

func test[T IN](x T) {

}

type IN interface {
	string | int | inter
}

type inter interface {
	String() string
}
