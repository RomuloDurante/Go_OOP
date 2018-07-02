package objectoriented

import "fmt"

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float64
}

//function constructor

func CreateCreditCard(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard{

	return &CreditCard{
		ownerName:ownerName,
		cardNumber:cardNumber,
		expirationMonth:expirationMonth,
		expirationYear:expirationYear,
		securityCode:securityCode,
		availableCredit: 5000, // call from some Api
	}
}

//Methods
func (c *CreditCard) ProcessPayment(amount float64) bool {
	if c.availableCredit > amount{
		fmt.Println("Payment done")
		c.availableCredit -= amount
		return true
	} else {
		fmt.Println("No enough money")
		return false
	}
}

// access to hide field
func (c *CreditCard) AvailableCredit() float64 {
	return c.availableCredit
}
