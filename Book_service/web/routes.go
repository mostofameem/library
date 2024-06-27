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
	mux.Handle(
		"POST /borrow",
		manager.With(
			http.HandlerFunc(handlers.BorrowBook),
			//middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"GET /profile",
		manager.With(
			http.HandlerFunc(handlers.Profile),
		),
	)
	mux.Handle(
		"PATCH /approve-request",
		manager.With(
			http.HandlerFunc(handlers.ApproveBorrowRequest),
		),
	)
	mux.Handle(
		"DELETE /reject-request",
		manager.With(
			http.HandlerFunc(handlers.RejectBorrowRequest),
		),
	)
	mux.Handle(
		"GET /get-borrow-request",
		manager.With(
			http.HandlerFunc(handlers.GetBorrowRequest),
		),
	)
}
