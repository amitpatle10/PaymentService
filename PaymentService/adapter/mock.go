package adapter

import (
	"fmt"

	"github.com/amitpatle/paymentservice/provider"
)

// MockProvider is a fake payment provider used for testing.
// It never calls any real API — just logs what it would do.
type MockProvider struct{}

// Charge pretends to charge the user and always succeeds.
func (m MockProvider) Charge(req provider.ChargeRequest) error {
	fmt.Printf("[MOCK] Charging %.2f %s for user %s (idempotency key: %s)\n",
		req.Payment.Amount,
		req.Payment.Currency,
		req.Payment.UserID,
		req.IdempotencyKey,
	)
	return nil // nil means success — no error
}

// Refund pretends to refund a payment and always succeeds.
func (m MockProvider) Refund(paymentID string) error {
	fmt.Printf("[MOCK] Refunding payment %s\n", paymentID)
	return nil
}
