package main

import (
	"github.com/blocktop/mp-common/config"
	"github.com/blocktop/mp-common/server"
	"github.com/blocktop/mp-common/server/middleware"
	chim "github.com/go-chi/chi/middleware"
	"github.com/stellar/go/support/http"
	"os"
	"time"
)

func main() {
	cfg := config.GetConfig()

	r := http.NewAPIMux(false)
	r.Use(middleware.HealthMiddleware)
	r.Use(middleware.JWTAuthMiddleWare)
	r.Use(chim.Timeout(time.Duration(cfg.HTTPServerRequestTimeout) * time.Second))

	setRoutes(r)

	server.RunHTTPServer(r)

	os.Exit(0)
}
