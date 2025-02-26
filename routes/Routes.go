package routes

import (
	"golang-api-lite/auth"
	"golang-api-lite/routes/handlers"
	"golang-api-lite/routes/pageHandlers"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Routes
func AddRoutes() *http.ServeMux {
	//config := utility.GetConfig()

	// First create mux
	mux := http.NewServeMux()

	// Swagger
	mux.HandleFunc("/swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	// API
	mux.HandleFunc("/foobar", auth.BasicAuth(handlers.HandleFooBar, "username", "password"))

	// Pages
	mux.HandleFunc("/home", pageHandlers.HandleHome)
	mux.HandleFunc("/component", pageHandlers.HandleComponent)

	return mux
}
