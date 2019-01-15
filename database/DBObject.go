package database

//Account struct in blocks
type Account struct {
	Address    string
	PublicKey  string
	Balance    float64
	Permission string
	Sequence   int64
	Code       string
	ID         int
}

type dbEngine interface {
	Connect() error
	Disconnect()

	InsertAccount(acc *Account) error

	UpdateAccount(id int, acc *Account) error

	GetAccount(id int) (*Account, error)
}
