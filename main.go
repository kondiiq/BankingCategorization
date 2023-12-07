package main

import (
	"banplication/db"
	"banplication/srv/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	Router := gin.Default()
	receiptRoute := Router.Group("/Receipt")
	receiptRoute.GET("", handlers.GetReceipts)
	receiptRoute.GET(":id", handlers.GetReceiptByID)
	receiptRoute.POST("", handlers.CreateReceipts)
	receiptRoute.PATCH(":id", handlers.UpdateReceipts)
	receiptRoute.DELETE(":id", handlers.DeleteReceipts)

	db.Connect()

	serverConfig := &http.Server{
		Addr:         os.Getenv("SERVER_PORT"),
		Handler:      Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	errSrvConfig := serverConfig.ListenAndServe()
	if errSrvConfig != nil {
		panic(errSrvConfig)
	}
}
