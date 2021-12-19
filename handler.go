package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

//Wrap API handler and returns standard http.HandlerFunc function
func Wrap[RQ any, RS any](handler func(ctx context.Context, request *RQ) (RS, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(RQ)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Fail to parse request body: %s", err.Error())))
			return
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
