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
		"POST /login",
		manager.With(
			http.HandlerFunc(handlers.Login),
			//middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"PATCH /update_user",
		manager.With(
			http.HandlerFunc(handlers.UpdateUser),
		),
	)
	mux.Handle(
		"GET /get_user",
		manager.With(
			http.HandlerFunc(handlers.GetUser),
		),
	)
	mux.Handle(
		"PATCH /approve-user",
		manager.With(
			http.HandlerFunc(handlers.ApproveUserRequest),
		),
	)
	mux.Handle(
		"DELETE /reject-user",
		manager.With(
			http.HandlerFunc(handlers.RejectUserRequest),
		),
	)

}
