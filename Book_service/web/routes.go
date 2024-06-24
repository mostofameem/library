package web

import (
	"book_service/web/handlers"
	"book_service/web/middlewares"
	"net/http"
)

func InitRouts(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /addbook",
		manager.With(
			http.HandlerFunc(handlers.Addbook),
		),
	)

	mux.Handle(
		"GET /showbooks",
		manager.With(
			http.HandlerFunc(handlers.GetBooks),
			//middlewares.AuthenticateJWT,
		),
	)

}
