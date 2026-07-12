package main

import (
	"log"
	"net/http"
	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8001")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", enableCORS(handleTripPreview))
	mux.HandleFunc("POST /trip/start", enableCORS(handleTripStart))
	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}
	
	mux.HandleFunc("/ws/riders", handleRidersWebSocket)
	mux.HandleFunc("/ws/drivers", handleDriversWebSocket)
	if err := server.ListenAndServe(); err != nil {

		log.Printf("HTTP server error: %v", err)
	}

}
