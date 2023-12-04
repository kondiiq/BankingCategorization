package handlers

import (
	"banplication/model"
	"github.com/gin-gonic/gin"
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

func GetReceipts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": receipts,
	})
}

func GetReceiptByID(c *gin.Context) {
	id := c.Param("id")
	for _, idx := range receipts {
		if idx.ID == id {
			c.IndentedJSON(http.StatusOK, idx)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not founded",
	})
	return
}

func CreateReceipts(c *gin.Context) {
	var newReceipt model.Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}
	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}

func UpdateReceipts(c *gin.Context) {
	var newReceipt model.Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}
	id := c.Param("id")
	for _, idx := range receipts {
		if idx.ID == id {
			idx.Sum = newReceipt.Sum
			idx.Address = newReceipt.Address
			idx.Currency = newReceipt.Currency
			idx.PaymentType = newReceipt.PaymentType
			idx.TransactionDate = newReceipt.TransactionDate
			idx.UpdatedAt = newReceipt.UpdatedAt
			c.IndentedJSON(http.StatusOK, newReceipt)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not founded",
	})
	return
}

func DeleteReceipts(c *gin.Context) {
	id := c.Param("id")
	if StringToUint(id) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
	}
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Deleted receipt",
	})
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
