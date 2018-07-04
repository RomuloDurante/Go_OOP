package main

import (
	"fmt"
	"github.com/RomuloDurante/GoSandBox/oopGo/messagePassing/channel"
)

func main() {
	chargeCh := make(chan float64)
	testChan := channel.CreateCreditAccount(chargeCh)
	chargeCh <- 500


	fmt.Println(testChan, chargeCh)

	//var a string
	//fmt.Scanln(&a)
}
