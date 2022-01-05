package api

import (
	"context"
	"encoding/json"
	"net/http"
)

//Wrap API handler and returns standard http.HandlerFunc function
func Wrap[RQ any, RS any](handler func(ctx context.Context, request *RQ) (RS, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(RQ)
		richifyRequest(req, r)
		switch r.Method {
		case http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodPut:
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
		}
		resp, err := handler(r.Context(), req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}
}

func richifyRequest[RQ any](req *RQ, baseRequest *http.Request) {
	if v, ok := (any)(req).(WithHeader); ok {
		v.WithHeader(baseRequest.Header)
	}
	if v, ok := (any)(req).(WithMethod); ok {
		v.WithMethod(baseRequest.Method)
	}
}

type WithHeader interface {
	WithHeader(header http.Header)
}

type WithMethod interface {
	WithMethod(method string)
}
