package payment

import "fmt"

type PaymentOption interface {
	ProcessPayment(float64) bool
}

type CreditAccount struct {
	accountNumber string
	accountOwner string
}

// constructor
func CreateAccount(accountNumber, accountOwner string)  *CreditAccount{
	return &CreditAccount{
		accountNumber:accountNumber,
		accountOwner:accountOwner,
	}
}

// Methods

func (c CreditAccount) ProcessPayment(amount float64) bool  {
	fmt.Println("it's work")
	return true
}
