package main

import (
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Wallet API
//	@version		1.0
//	@description	Sophisticated Wallet API
//	@host			localhost:1323
func main() {
	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(p)
	e.GET("/api/v1/wallets", handler.WalletHandler) // complete the code
	e.POST("/api/v1/wallets", handler.CreateWalletHandler) // complete the code
	e.PUT("/api/v1/wallets", handler.UpdateWalletHandler) // complete the code
	e.DELETE("/api/v1/wallets/:id", handler.DeleteWalletHandler) // complete the code
	e.GET("/users/:id/wallets", handler.GetWalletSpecificByUserIdHandler) // complete the code
	e.Logger.Fatal(e.Start(":1323"))
}
