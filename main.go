package main

import (
	"fmt"
	"log"

	"github.com/JunYeong-dev/BankAndDictionary/accounts"
)

func main() {
	account := accounts.NewAccount("judy")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
		// fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
