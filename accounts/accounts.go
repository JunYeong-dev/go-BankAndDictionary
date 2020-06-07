package accounts

import (
	"errors"
	"fmt"
)

// Account struct - private와 public은 대소문자로 구분(소문자: private, 대문자: public)
type Account struct {
	owner   string
	balance int
}

// error를 미리 변수로 정의
var errNoMoney = errors.New("Can`t withdraw")

// NewAccount - Go에서는 constructor가 없기 때문에 function으로 construct하거나 struct를 만들도록 함
// struct 인스턴스(*Account)를 생성한 후 pointer-주소(&account)를 return하는 방식으로 constructor역할을 함
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit - method, Account의 balance 값을 변경(+)
// (a Account) : receiver - a는 이름이고 type은 Account
// 이름을 짓는 규칙은 struct의 첫 글자를 따서 '소문자'로 지어야 함
// 기본적으로 Go의 경우 (a Account)의 형태로 recevier로 하는 경우 복사본을 만들게 됨
// main.go 쪽의 account에 접근 하는 것이 아닌 복사본의 amount의 값을 변경하게 되는 것
// 그렇기 때문에 직접 값을 변경 하기 위해선 pointer로 접근하는 것이 필요함
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance - Account의 balance 값을 return
func (a Account) Balance() int {
	return a.balance
}

// Withdraw - method, Account의 balance 값을 변경(-)
// Go에서는 try - catch 가 없기 때문에 직접 error를 return 하고 체크해줘야 한다
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	// nil - error의 return타입 중 하나로서 다른 언어의 null이나 none과 같음
	return nil
}

// ChangeOwner - Account의 owner 값을 변경
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner - Account의 owner 값을 return
func (a Account) Owner() string {
	return a.owner
}

// Java의 toString()함수와 같다고 생각하면 됨
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
