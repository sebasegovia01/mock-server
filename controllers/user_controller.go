package controllers

import (
	"mock-server/models"
	"mock-server/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userData []models.User

func InitializeData(data []models.User) {
	userData = data
}

func GetPersonalIdentification(c *gin.Context) {
	customerID := c.Param("customerIdentification")
	if customerID == "" {
		logger.ErrorLogger.Println("customerIdentification not provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "customerIdentification is required"})
		return
	}

	logger.InfoLogger.Printf("Searching for user with ID: %s", customerID)
	for _, user := range userData {
		if user.PersonalIdentification.CustomerIdentification == customerID {
			logger.InfoLogger.Printf("User found: %s %s",
				user.PersonalIdentification.CustomerFirstName,
				user.PersonalIdentification.CustomerLastName)
			c.JSON(http.StatusOK, user.PersonalIdentification)
			return
		}
	}

	logger.ErrorLogger.Printf("User not found for ID: %s", customerID)
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func GetCustomerAccounts(c *gin.Context) {
	customerID := c.Param("customerIdentification")
	if customerID == "" {
		logger.ErrorLogger.Println("customerIdentification not provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "customerIdentification is required"})
		return
	}

	logger.InfoLogger.Printf("Searching for accounts for user with ID: %s", customerID)
	for _, user := range userData {
		if user.PersonalIdentification.CustomerIdentification == customerID {
			logger.InfoLogger.Printf("Account found for user: %s %s",
				user.PersonalIdentification.CustomerFirstName,
				user.PersonalIdentification.CustomerLastName)
			c.JSON(http.StatusOK, user.Accounts.AccountInformation)
			return
		}
	}

	logger.ErrorLogger.Printf("No accounts found for user with ID: %s", customerID)
	c.JSON(http.StatusNotFound, gin.H{"error": "accounts not found for the specified customer"})
}

func GetAccountBalance(c *gin.Context) {
	productID := c.Param("productIdentification")

	for _, user := range userData {
		if user.Accounts.AccountInformation.ProductIdentification == productID {
			c.JSON(http.StatusOK, user.Accounts.Balances)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
}
