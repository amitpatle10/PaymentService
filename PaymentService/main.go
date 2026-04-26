package main

import (
	"fmt"
	"time"

	"github.com/amitpatle/paymentservice/adapter"
	"github.com/amitpatle/paymentservice/models"
	"github.com/amitpatle/paymentservice/provider"
)

// processPayment accepts any PaymentProvider — Mock, Stripe, Razorpay, anything.
// It doesn't know or care which one it gets. That's the power of the interface.
func processPayment(p provider.PaymentProvider, req provider.ChargeRequest) {
	err := p.Charge(req)
	if err != nil {
		fmt.Printf("Payment failed: %s\n", err)
		return
	}
	fmt.Println("Payment successful!")
}

func main() {
	fmt.Println("PaymentService starting...")

	p := models.Payment{
		ID:        "pay_001",
		UserID:    "user_42",
		Amount:    999.99,
		Currency:  "INR",
		Provider:  "stripe",
		Status:    models.StatusPending,
		CreatedAt: time.Now(),
	}

	req := provider.ChargeRequest{
		Payment:        p,
		IdempotencyKey: "idem_key_001",
	}

	fmt.Println("--- Using Mock Provider ---")
	processPayment(adapter.MockProvider{}, req)

	fmt.Println("--- Using Stripe Provider ---")
	processPayment(adapter.StripeProvider{APIKey: "sk_test_abc123"}, req)
}
