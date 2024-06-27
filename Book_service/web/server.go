package web

import (
	"book_service/config"
	"book_service/web/middlewares"
	"book_service/web/swagger"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
)

func StartServer(wg *sync.WaitGroup) {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()

	InitRouts(mux, manager)

	// Enable CORS on the mux
	handler := middlewares.EnableCors(mux)

	swagger.SetupSwagger(mux, manager)

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
