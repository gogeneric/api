package main

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/gogeneric/api"
)

var tpl *template.Template

func main() {
	h := &http.Server{Addr: "0.0.0.0:3000"}
	mux := http.NewServeMux()
	h.Handler = mux
	var err error
	tpl, err = template.ParseGlob("./*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// Here is magic!
	mux.Handle("/hello", api.Wrap(handleHello))

	if err := h.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func handleHello(ctx context.Context, req *helloRequest) (*helloResponse, error) {
	return &helloResponse{
		Message:  "Hello, " + req.Name,
		template: "tpl.gohtml",
	}, nil
}

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	template string
	Message  string
}

func (r *helloResponse) Render() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := tpl.ExecuteTemplate(buf, r.template, r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
