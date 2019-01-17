package database

import (
	hsBC "github.com/gallactic/hubble_service/blockchain"
)

//Adapter for data base
type Adapter interface {
	Connect() error
	Disconnect() error

	//Account Handling
	InsertAccount(acc *hsBC.Account) error
	UpdateAccount(id int, acc *hsBC.Account) error
	GetAccount(id int) (*hsBC.Account, error)

	//Blocks Handling
	InsertBlock(acc *hsBC.Account) error
	UpdateBlock(id int, acc *hsBC.Account) error
	GetBlock(id int) (*hsBC.Account, error)
	GetBlocksTableLastID() (uint64, error)
}
