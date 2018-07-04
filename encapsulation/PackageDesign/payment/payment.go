package payment

type CreditAccount struct {
	accountNumber string
	accountOwner string

}

// constructor
func CreateAccount(accountNumber, accountOwner string) *CreditAccount {
	return &CreditAccount{
		accountNumber: accountNumber,
		accountOwner: accountOwner,
	}
}

//Methods

// Account number ..
func (c CreditAccount) AccountNumber() string  {

	return c.accountNumber
}

// Account Owner..
func (c CreditAccount) AccountOwner() string  {

	return c.accountOwner
}