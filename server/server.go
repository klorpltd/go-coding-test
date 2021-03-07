package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Run(ctx context.Context, port int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", userHandler())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Printf("http server shutting down\n")
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Printf("http server shutdown err: %v\n", err)
				return
			}
		}
	}(ctx)

	return srv.ListenAndServe()
}

func userHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: implement me
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}
