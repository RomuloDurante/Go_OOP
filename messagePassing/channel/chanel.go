package channel

import "fmt"

type CreditAccount struct {}

func (c *CreditAccount) processPayment(amount float64)  {
	fmt.Println("Process credit card payment...")
}

//func constructor with channel
func CreateCreditAccount(chargeCh chan float64) *CreditAccount  {
	//empty struct
	creditAccount := &CreditAccount{}

	go func(chargeCh chan float64) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}