package main

import (
	"net/http"
	"os"

	"github.com/fajarfadilla/mokus/internal/components"
	"github.com/rs/zerolog"
)

func main() {
	mux := http.NewServeMux()
	log := zerolog.New(os.Stdout)

	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset utf-8")
		components.Hello("fajarfadilla").Render(r.Context(), w)
	})

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	log.Info().Msg("server started")
	err := http.ListenAndServe(s.Addr, s.Handler)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
		return
	}
	log.Info().Msg("server stopped")
}
