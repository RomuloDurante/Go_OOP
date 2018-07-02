package objectoriented

import "fmt"

type Cash struct{}

func CreateCashAccount() *Cash  {
	return &Cash{}
}

func (c Cash) ProcessPayment(amount float64) bool {
	fmt.Println("Payment done")
	return true
}
