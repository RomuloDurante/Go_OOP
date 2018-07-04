package main

import "fmt"

type Account struct{}

func (a *Account) Funds() float64  {

	fmt.Println("Listing avaliable funds")

	return 10
}
//**********************
type Account2 struct{}

func (a *Account2) Funds() float64  {

	fmt.Println("Listing avaliable funds")

	return 20
}

/// credit type
type CreditAccount struct {
	Account //composition
	Account2 // 

}

func (a *CreditAccount) Funds() float64  {

	return a.Account.Funds() + a.Account2.Funds()
}


func main()  {
	ca := &CreditAccount{}

	ca.Funds()
}