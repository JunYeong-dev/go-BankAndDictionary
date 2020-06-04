package accounts

// Account struct - private와 public은 대소문자로 구분(소문자: private, 대문자: public)
type Account struct {
	owner   string
	balance int
}

// NewAccount - Go에서는 constructor가 없기 때문에 function으로 construct하거나 struct를 만들도록 함
// struct 인스턴스(*Account)를 생성한 후 pointer-주소(&account)를 return하는 방식으로 constructor역할을 함
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
