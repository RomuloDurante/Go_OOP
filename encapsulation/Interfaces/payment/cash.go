package payment

import "fmt"

type Cash struct {}

func CreateCash() *Cash {
	return &Cash{}
}


func (c Cash) ProcessPayment(amount float64) bool  {
	fmt.Println("Process a cash...")
	return true
}
