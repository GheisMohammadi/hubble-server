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

//Account struct in blocks
type Block struct {
	ID             int
	Height         int
	Block_Hash     string
	Pre_Block_Hash string
	Time_Stamp     string
	Txn_Counts     int
}

//Adapter for data base
type Adapter interface {
	Connect() error
	Disconnect()

	//Account Handling
	InsertAccount(acc *Account) error
	UpdateAccount(id int, acc *Account) error
	GetAccount(id int) (*Account, error)

	//Blocks Handling
	InsertBlock(acc *Account) error
	UpdateBlock(id int, acc *Account) error
	GetBlock(id int) (*Account, error)
	GetBlocksTableLastID() (uint64, error)
}
