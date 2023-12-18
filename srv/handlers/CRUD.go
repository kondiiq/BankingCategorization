package handlers

import (
	"banplication/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var receipts = []model.Receipt{
	{ID: "1", Sum: 21.37, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "2", Sum: 21.73, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "3", Sum: 73.21, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "4", Sum: 12.37, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "5", Sum: 12.73, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "6", Sum: 73.12, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "7", Sum: 37.21, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "8", Sum: 37.12, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "9", Sum: 12.37, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
	{ID: "10", Sum: 13.27, Address: "Represent street 215, London, Great Britain", PaymentType: "Cash", Currency: "GBF"},
}

var db *gorm.DB

func GetReceipts(c *gin.Context) {
	var receipts []model.Receipt
	if err := db.Find(&receipts).Error; err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, receipts)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": receipts,
	//})
}

func GetReceiptByID(c *gin.Context) {
	id := c.Param("id")
	var receipt model.Receipt
	if err := db.Where("id= ?", id).First(&receipt); err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
		return
	} else {
		c.JSON(200, receipt)
	}
	//for _, idx := range receipts {
	//	if idx.ID == id {
	//		c.IndentedJSON(http.StatusOK, idx)
	//		return
	//	}
	//}
	//c.JSON(http.StatusNotFound, gin.H{
	//	"code":    404,
	//	"message": "Not found",
	//})
	//return
}

func CreateReceipts(c *gin.Context) {
	var newReceipt model.Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		log.Fatal(err)
		return
	}
	db.Create(&newReceipt)
	//receipts = append(receipts, newReceipt)
	c.JSON(http.StatusCreated, newReceipt)
}

func UpdateReceipts(c *gin.Context) {
	id := c.Param("id")
	var updateBody model.Receipt
	if err := db.Where("id = ?", id).First(&updateBody); err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
		return
	}
	c.BindJSON(&updateBody)
	db.Save(&updateBody)
	c.JSON(200, updateBody)
	//if err := c.BindJSON(&updateBody); err != nil {
	//	c.IndentedJSON(http.StatusBadRequest, err)
	//	return
	//}
	//
	//for _, idx := range receipts {
	//	if idx.ID == id {
	//		idx.Sum = updateBody.Sum
	//		idx.Address = updateBody.Address
	//		idx.PaymentType = updateBody.PaymentType
	//		idx.TransactionDate = updateBody.TransactionDate
	//		idx.Currency = updateBody.Currency
	//		c.JSON(http.StatusOK, &updateBody)
	//		return
	//	}
}

func DeleteReceipts(c *gin.Context) {
	id := c.Param("id")
	var receipt model.Receipt
	deleteQuery := db.Where("id = ?", id).Delete(&receipt)
	fmt.Println(deleteQuery)
	c.JSON(http.StatusOK, gin.H{
		"id #" + id: "deleted"})
	//if StringToUint(id) < 1 {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Bad request",
	//	})
	//}
	//c.JSON(http.StatusNoContent, gin.H{
	//	"message": "Deleted receipt",
	//})
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
