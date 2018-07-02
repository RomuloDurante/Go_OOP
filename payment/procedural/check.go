package procedural

type ChekingAccount struct {
	AccountOwner string
	RoutingNumber string
	AccountNumber string
	Balance float64
}

func PayWithCheck(account *ChekingAccount, amount float64) bool  {

	if account.Balance > amount {
		account.Balance -= amount
		return true
	} else {
		return false
	}

}

