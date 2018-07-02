package procedural

type CreditCard struct {
	OwnerName       string
	CardNumber      string
	ExpirationMonth int
	ExpirationYear  int
	SecurityCode    int
	AvailableCredit float64
}

func PayWithCredit(card *CreditCard, amount float64) bool {
	if card.AvailableCredit > amount {
		card.AvailableCredit -= amount
		return true
	} else {
		return false
	}
}
