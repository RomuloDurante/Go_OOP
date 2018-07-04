package main

import (
	"fmt"
	"github.com/RomuloDurante/GoSandBox/oopGo/encapsulation/PackageDesign/payment"
)

func main() {

	fromPackage := payment.CreateAccount("123456", "Romulo")
	fmt.Println(fromPackage.AccountNumber(), fromPackage.AccountOwner())
}
