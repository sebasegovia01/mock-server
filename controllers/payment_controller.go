package controllers

import (
	"fmt"
	"math/rand"
	"mock-server/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var paymentReq models.PaymentRequest
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate unique TransactionId
	transactionId := generateTransactionId()

	// Determine payment status based on probability
	status := getRandomStatus(70) // 70% success rate

	// Adjust HTTP status code based on result
	httpStatus := http.StatusOK
	// if status == models.Failed {
	// 	httpStatus = http.StatusInternalServerError
	// }

	response := models.PaymentResponse{
		Amount:          paymentReq.Amount,
		Currency:        paymentReq.Currency,
		SourceAccount:   paymentReq.SourceAccount,
		DestAccount:     paymentReq.DestAccount,
		Description:     paymentReq.Description,
		PaymentType:     paymentReq.PaymentType,
		TransactionId:   transactionId,
		PaymentDate:     time.Now(),
		BeneficiaryName: paymentReq.BeneficiaryName,
		Status:          status,
	}

	c.JSON(httpStatus, response)
}

// Function to generate unique TransactionId
func generateTransactionId() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	random := rand.Intn(10000)
	return fmt.Sprintf("TRX%d%04d", timestamp, random)
}

// Function to determine random status based on success rate
func getRandomStatus(successRate int) models.PaymentStatus {
	// Generate random number between 1 and 100
	randomNum := rand.Intn(100) + 1

	if randomNum <= successRate {
		return models.Completed
	}
	return models.Failed
}
