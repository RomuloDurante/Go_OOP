package objectoriented

type PaymentOption interface {
	ProcessPayment(float64) bool
} 
