package database

type dbEngine interface {
	Connect() error
	Disconnect()

	InsertAccount(address string, publicKey string, balance float64,
		perm string, seq string, code string) error
}
