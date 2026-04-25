package main

import (
	"fmt"
	"time"

	"github.com/amitpatle/paymentservice/models"
)

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

	fmt.Printf("Payment ID: %s\n", p.ID)
	fmt.Printf("Amount: %.2f %s\n", p.Amount, p.Currency)
	fmt.Printf("Status: %s\n", p.Status)
	fmt.Printf("Created: %s\n", p.CreatedAt.Format("2006-01-02 15:04:05"))
}
