package web

import (
	"net/http"
	"user_service/web/handlers"
	"user_service/web/middlewares"
)

func InitRouts(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /register",
		manager.With(
			http.HandlerFunc(handlers.Create),
		),
	)
	mux.Handle(
		"PATCH /approved",
		manager.With(
			http.HandlerFunc(handlers.Approved),
		),
	)

	mux.Handle(
		"POST /login",
		manager.With(
			http.HandlerFunc(handlers.Login),
			//middlewares.AuthenticateJWT,
		),
	)
}
