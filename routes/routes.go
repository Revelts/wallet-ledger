package routes

import (
	"github.com/go-chi/chi/v5"
	"go-testing/controller"
	"log"
	"net/http"
)

func InitRoutes() {
	r := chi.NewRouter()

	r.Post("/deposit", controller.DepositHandler)
	r.Post("/withdraw", controller.WithdrawHandler)
	r.Post("/transfer", controller.TransferHandler)
	r.Get("/wallet/balance", controller.WalletBalanceHandler)
	r.Get("/wallet/history", controller.WalletHistoryHandler)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
