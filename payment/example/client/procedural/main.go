package main

import (
	"fmt"
	"strings"
	"github.com/RomuloDurante/GoSandBox/oopGo/payment/procedural"
)

func main() {
	const amount = 500

	// with cash payment
	fmt.Println("Payment with cash")
	procedural.PayWithCash(amount)
	fmt.Println(strings.Repeat("*", 10) + "\n\n")

	//with credicard
	credit := &procedural.CreditCard{
		OwnerName:       "Romulo",
		CardNumber:      "1111-2222-3333-4444",
		ExpirationMonth: 5,
		ExpirationYear:  2021,
		SecurityCode:    123,
		AvailableCredit: 5000,
	}

	fmt.Println("Payment with Credit")
	fmt.Printf("Initial balance: $%.2f\n", credit.AvailableCredit)
	procedural.PayWithCredit(credit, amount)
	fmt.Printf("Balance Now: $%.2f\n", credit.AvailableCredit)
	fmt.Println(strings.Repeat("*", 10) + "\n\n")

	//with check
	checking := &procedural.ChekingAccount{
		AccountOwner: "Romulo",
		RoutingNumber: "01010101",
		AccountNumber: "0123654789",
		Balance: 250,
	}

	fmt.Println("Payment with Check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance)
	procedural.PayWithCheck(checking, amount)
	fmt.Printf("Balance Now: $%.2f\n", checking.Balance)
	fmt.Println(strings.Repeat("*", 10) + "\n\n")

	fmt.Print("Not enough money \n\n")
	checking.Balance = 5000

	fmt.Println("Payment with Check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance)
	procedural.PayWithCheck(checking, amount)
	fmt.Printf("Balance Now: $%.2f\n", checking.Balance)
	fmt.Println(strings.Repeat("*", 10) + "\n\n")
}
