package models

import "time"

// PaymentStatus is a custom type based on string.
// Instead of using raw strings like "pending", "success" everywhere,
// we define named constants — no typos, no guessing.
type PaymentStatus string

const (
	StatusPending  PaymentStatus = "pending"
	StatusSuccess  PaymentStatus = "success"
	StatusFailed   PaymentStatus = "failed"
)

// Payment represents a single payment transaction.
// This is the core data structure of our entire service.
type Payment struct {
	ID        string        // unique identifier for this payment
	UserID    string        // who is making the payment
	Amount    float64       // how much (e.g. 999.99)
	Currency  string        // "INR", "USD" etc.
	Provider  string        // "stripe", "razorpay", "paypal"
	Status    PaymentStatus // current state of the payment
	CreatedAt time.Time     // when was this payment created
}
