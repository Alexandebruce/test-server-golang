package handlers

import (
	"github.com/Alexandebruce/test-server-golang/db/logging"
	"io"
	"log"
	"net/http"
	"time"
)

func PingHandler(logger *logging.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "I'm test-server-golang microservice!"); err != nil {
			log.Printf("error to write response: %s", err)
		}

		if err := logger.Insert(logging.PingLogMessage{
			Message: "Ping",
			Date: time.Now(),
		}); err != nil {
			log.Printf("error to write log message: %s", err)
		}
	})
}