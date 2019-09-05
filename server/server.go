package server

import (
	"goLang-jwt-auth/server/middleware"
	"log"
	"net/http"
)

func StartServer(hostname string, port string) error {
	host := hostname + ":" + port

	log.Printf("Listening on: %s", host)

	handler := middleware.NewHandler()

	http.Handle("/", handler)
	return http.ListenAndServe(host, nil)
}