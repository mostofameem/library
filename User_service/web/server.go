package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"user_service/config"
	"user_service/web/middlewares"
	"user_service/web/swagger"
)

// StartServer initializes and starts the HTTP server
func StartServer(wg *sync.WaitGroup) {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()

	InitRouts(mux, manager)
	swagger.SetupSwagger(mux, manager)

	handler := middlewares.EnableCors(mux)

	wg.Add(1)

	go func() {
		defer wg.Done()

		conf := config.GetConfig()
		addr := fmt.Sprintf(":%d", conf.HttpPort)
		slog.Info(fmt.Sprintf("Listening at %s", addr))

		if err := http.ListenAndServe(addr, handler); err != nil {
			slog.Error(err.Error())
		}
	}()
}
