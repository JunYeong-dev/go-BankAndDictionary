package main

import (
	"fmt"

	"github.com/JunYeong-dev/BankAndDictionary/accounts"
)

func main() {
	account := accounts.NewAccount("judy")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
}
