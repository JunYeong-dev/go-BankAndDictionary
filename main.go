package main

import (
	"fmt"

	"github.com/JunYeong-dev/BankAndDictionary/accounts"
)

func main() {
	account := accounts.NewAccount("judy")
	fmt.Println(account)
}
