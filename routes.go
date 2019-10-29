package main

import (
	"github.com/blocktop/mp-common/server/middleware"
	"github.com/blocktop/mp-transfer-server/handlers"
	"github.com/go-chi/chi"
)

func setRoutes(r *chi.Mux) {
	r.Get("/health", middleware.HealthHandler)

//	r.Get("/deposit", handleGetDeposit)
//	r.Get("/withdraw", handleGetWithdraw)
	r.Get("/info", handlers.HandleGetInfo)
//	r.Get("/transactions", handleGetTransactions)
//	r.Get("/transaction", handleGetTransaction)

}
