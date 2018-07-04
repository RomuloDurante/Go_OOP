package main

import (
	"github.com/RomuloDurante/GoSandBox/oopGo/encapsulation/Interfaces/payment"
)

func main() {

	var option payment.PaymentOption

	option = payment.CreateAccount("123456", "Romulo")
	option.ProcessPayment(500)

	option = payment.CreateCash()
	option.ProcessPayment(500)




}
