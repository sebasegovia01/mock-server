package models

import "time"

// PaymentRequest es la estructura para el payload de entrada
type PaymentRequest struct {
	Amount          float64     `json:"amount" binding:"required"`
	Currency        string      `json:"currency" binding:"required"`
	SourceAccount   string      `json:"sourceAccount" binding:"required"`
	DestAccount     string      `json:"destAccount" binding:"required"`
	Description     string      `json:"description" binding:"required"`
	PaymentType     PaymentType `json:"paymentType" binding:"required"`
	BeneficiaryName string      `json:"beneficiaryName" binding:"required"`
}

// PaymentResponse es la estructura para la respuesta
type PaymentResponse struct {
	Amount          float64       `json:"amount"`
	Currency        string        `json:"currency"`
	SourceAccount   string        `json:"sourceAccount"`
	DestAccount     string        `json:"destAccount"`
	Description     string        `json:"description"`
	PaymentType     PaymentType   `json:"paymentType"`
	TransactionId   string        `json:"transactionId"`
	PaymentDate     time.Time     `json:"paymentDate"`
	BeneficiaryName string        `json:"beneficiaryName"`
	Status          PaymentStatus `json:"status"`
}
