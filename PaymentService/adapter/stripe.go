package adapter

import (
	"fmt"

	"github.com/amitpatle/paymentservice/provider"
)

// StripeProvider is the Stripe implementation of PaymentProvider.
// In production, this is where real Stripe API calls would go.
type StripeProvider struct {
	APIKey string // Stripe secret key
}

// Charge sends a charge request to Stripe.
func (s StripeProvider) Charge(req provider.ChargeRequest) error {
	// In production: call Stripe's API here using s.APIKey
	// For now, we simulate it
	fmt.Printf("[STRIPE] Charging %.2f %s for user %s (idempotency key: %s)\n",
		req.Payment.Amount,
		req.Payment.Currency,
		req.Payment.UserID,
		req.IdempotencyKey,
	)
	return nil
}

// Refund sends a refund request to Stripe.
func (s StripeProvider) Refund(paymentID string) error {
	// In production: call Stripe's refund API here
	fmt.Printf("[STRIPE] Refunding payment %s\n", paymentID)
	return nil
}
