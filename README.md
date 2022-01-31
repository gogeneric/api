# API wrapper

API handler wrapper

## Usage

### api.Wrap(handler)

Function Wrap wraps API handler and returns standard http.HandlerFunc. It encapsulate body parsing.

#### Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/neonxp/api"
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

// Our API handler with custom request and response types
func handleHello(ctx context.Context, req *helloRequest) (*helloResponse, error) {
	return &helloResponse{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

// Custom request type
type helloRequest struct {
	Name string `json:"name"`
}

// Custom response type
type helloResponse struct {
	Message string `json:"message"`
}

```
