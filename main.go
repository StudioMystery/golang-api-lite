package main

import (
	"net/http"

	_ "golang-api-lite/docs" // swagger
	"golang-api-lite/routes"
	"golang-api-lite/utility"

	"github.com/rs/cors"
)

func main() {
	var config = utility.GetConfig()

	corsHandler := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedOrigins:   []string{"http://" + config.Foobar_FrontendHostname},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Authorization"},
		Debug:            true,
	})

	mux := routes.AddRoutes()
	handler := corsHandler.Handler(mux)

	if config.Env == "prod" {
		certFile := "./certs/foobar.pem"
		keyFile := "./certs/foobar.key"
		http.ListenAndServeTLS(":8090", certFile, keyFile, handler)
	} else {
		http.ListenAndServe(":8090", handler)
	}
}
