package main

import (
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

	oRouter := gin.Default()
	oReceipt := oRouter.Group("/Receipt")
	oReceipt.GET("", handlers.GetReceipts)
	oReceipt.GET(":id", handlers.GetReceiptByID)
	oReceipt.POST("", handlers.CreateReceipts)
	oReceipt.PATCH(":id", handlers.UpdateReceipts)
	oReceipt.DELETE(":id", handlers.DeleteReceipts)

	oServerConfig := &http.Server{
		Addr:         os.Getenv("SERVER_PORT"),
		Handler:      oRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	oServerConfig.ListenAndServe()
}
