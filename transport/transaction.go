package transport

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/bpjs-test-golang/data"
	"github.com/uchupx/bpjs-test-golang/data/model"
)

type Payload struct {
	RequestId uint64        `json:"request_id"`
	Data      []Transaction `json:"data"`
}

type Transaction struct {
	Id        uint64 `json:"id" db:"id"`
	Customer  string `json:"customer" db:"customer"`
	Quantity  uint64 `json:"quantity" db:"quantity"`
	Price     string `json:"price" db:"price"`
	Timestamp string `json:"timestamp" db:"timestamp"`
}

type TransactionHandler struct {
	TransactionRepository data.TransactionRepository
}

func (h TransactionHandler) Posts(c *gin.Context) {
	var req Payload

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	mapTransactions := make(map[uint][]model.Transaction)
	count := 0
	countIdx := 0

	for _, item := range req.Data {
		if count > 999 {
			countIdx++
			count = 0
		}

		date, err := time.Parse("2006-01-02 15:04:05", item.Timestamp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		mapTransactions[uint(countIdx)] = append(mapTransactions[uint(countIdx)], model.Transaction{
			Id:        item.Id,
			Customer:  item.Customer,
			Quantity:  item.Quantity,
			Price:     price,
			Timestamp: date,
		})

		count++
	}

	for _, transactions := range mapTransactions {
		_, err := h.TransactionRepository.Inserts(c, transactions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, "success")
	return
}

func (h TransactionHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
	return
}
