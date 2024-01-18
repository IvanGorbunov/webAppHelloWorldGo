package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request to /hello endpoint.")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello, world!")
	})

	slog.Info("Starting server on port 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		slog.Error("Application finished with an error", "error", err)
	}
}
