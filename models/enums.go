package models

type PaymentType string
type PaymentStatus string

const (
	Transfer PaymentType = "TRANSFER"
	Payment  PaymentType = "PAYMENT"
)

const (
	Pending   PaymentStatus = "PENDING"
	Completed PaymentStatus = "COMPLETED"
	Failed    PaymentStatus = "FAILED"
)

type ResultStatus string

const (
	OK      ResultStatus = "OK"
	ERROR   ResultStatus = "ERROR"
	WARNING ResultStatus = "WARNING"
)

type CanonicalErrorType string

const (
	NEG CanonicalErrorType = "NEG"
	TEC CanonicalErrorType = "TEC"
	SEG CanonicalErrorType = "SEG"
)
