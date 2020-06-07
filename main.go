package main

import (
	"fmt"

	"github.com/JunYeong-dev/BankAndDictionary/accounts"
)

func main() {
	account := accounts.NewAccount("judy")
	account.Deposit(10)
	// err := account.Withdraw(20)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	// fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())

	// String()함수가 정의 되어있다면 기본적으로 String()함수를 실행하게 됨
	fmt.Println(account)
}
