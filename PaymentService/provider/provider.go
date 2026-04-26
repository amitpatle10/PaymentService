package provider

import "github.com/amitpatle/paymentservice/models"

// ChargeRequest holds all the information needed to charge a user.
// Instead of passing 5 separate arguments to Charge(), we group them in a struct.
type ChargeRequest struct {
	Payment     models.Payment
	IdempotencyKey string // unique key to prevent double charges
}

// PaymentProvider is the contract every payment provider must follow.
// Any struct that has Charge() and Refund() methods automatically satisfies this interface.
type PaymentProvider interface {
	Charge(req ChargeRequest) error
	Refund(paymentID string) error
}
