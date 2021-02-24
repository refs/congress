package main

import (
	chttp "github.com/refs/congress/pkg/server/http"
	"net/http"
)

func main() {
	server := chttp.New()
	server.Address = "0.0.0.0:9337"
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello, micro service!")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	server.Handler = h
	if err := server.Run(); err != nil {
		panic(err)
	}
}
