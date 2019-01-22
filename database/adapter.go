package database

import (
	hsBC "github.com/gallactic/hubble_server/blockchain"
)

//Adapter for data base
type Adapter interface {
	Connect() error
	Disconnect() error

	//Account Handling
	InsertAccount(acc *hsBC.Account) error
	UpdateAccount(id int, acc *hsBC.Account) error
	GetAccount(id int) (*hsBC.Account, error)
	GetAccountsTableLastID() (uint64, error)

	//Blocks Handling
	InsertBlock(b *hsBC.Block) error
	UpdateBlock(id int, b *hsBC.Block) error
	GetBlock(id int) (*hsBC.Block, error)
	GetBlocksTableLastID() (uint64, error)

	//Transactions Handling
	InsertTx(b *hsBC.Block) error
	UpdateTx(id int, b *hsBC.Block) error
	GetTx(id int) (*hsBC.Block, error)
	GetTXsTableLastID() (uint64, error)

}
