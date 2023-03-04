package handler

import (
	"bpjs/config"
	"bpjs/middleware"
	"bpjs/pkg/myservice"
	"net/http"

	"github.com/eben-hk/confide"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

func Handler(l myservice.Service) http.Handler {

	r := mux.NewRouter()
	r = r.PathPrefix("/" + config.AppName + "/v1").Subrouter()

	// Public APIs
	r.HandleFunc("/health", getHealthStatus()).Methods("GET")

	// Protected APIs
	// Listing

	// Not found
	r.NotFoundHandler = r.NewRoute().HandlerFunc(notFound()).GetHandler()

	// Protected APIs
	// Adding API
	r.HandleFunc("/bulk", addOrderList(l)).Methods("POST")

	// Not allowed
	r.MethodNotAllowedHandler = r.NewRoute().HandlerFunc(notAllowed()).GetHandler()

	r.Use(middleware.InternalLogger())

	return r
}

// notFound to handle undefined route
func notFound() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		confide.JSON(w, confide.Payload{Code: confide.FCodeUriNotFound})
	}
}

// notAllowed to handle unallowed method
func notAllowed() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		confide.JSON(w, confide.Payload{Code: confide.FCodeMethodNotAllowed})
	}
}

// getWelcome to check health status of the server
func getHealthStatus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		confide.JSON(w, confide.Payload{Code: confide.SCodeUp, IsSuccess: true})
	}
}
