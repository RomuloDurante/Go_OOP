package objectoriented

import "fmt"

type CheckingAccount struct {
	accountOwner string
	routingNumber string
	accountNumber string
	balance float64
}

//function constructor
func CreateCheckingAccount(accountOwner, routingNumber, accountNumber string) * CheckingAccount {

	return &CheckingAccount{
		accountNumber:accountNumber,
		routingNumber:routingNumber,
		accountOwner:accountOwner,
		balance: 250, // call from a api
	}
}


//Methods
func (c *CheckingAccount) ProcessPayment(amount float64) bool {
	if c.balance > amount{
		fmt.Println("Payment done")
		c.balance -= amount
		return true
	} else {
		fmt.Println("No enough money")
		return false
	}

}


// access to hide field
func (c *CheckingAccount) Balance() float64 {
	return c.balance
}

