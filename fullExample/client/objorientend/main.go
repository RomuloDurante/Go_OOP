package main

import (
	"fmt"
	"strings"

	"github.com/RomuloDurante/GoSandBox/oopGo/fullExample/objectoriented"
	"github.com/RomuloDurante/GoSandBox/oopGo/encapsulation/Interfaces/payment"
)

func main() {
	const amount = 500

	//use the interface
	var option payment.PaymentOption


	// use cash payment *********************************************************
	fmt.Println("Payment with cash")
	option = payment.CreateCash() // example using interface
	option.ProcessPayment(amount)
	fmt.Println(strings.Repeat("*", 10) + "\n\n")

	// use credit ****************************************************************
	credit := objectoriented.CreateCreditCard( // here i'm using function constructor
		"Romulo",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	fmt.Println("Payment with credit")
	fmt.Printf("Initial balance: $%.2f\n", credit.AvailableCredit())
	credit.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f\n", credit.AvailableCredit())
	fmt.Println(strings.Repeat("*", 10) + "\n\n")

	// use check ******************************************************************
	check := objectoriented.CreateCheckingAccount( // here i'm using function constructor
		"Romulo",
		"123-6589",
		"254125")

	fmt.Println("Payment with check")
	fmt.Printf("Initial balance: $%.2f\n", check.Balance())
	check.ProcessPayment(amount)
	fmt.Printf("Balance now: $%.2f\n", check.Balance())
	fmt.Println(strings.Repeat("*", 10) + "\n\n")

}
