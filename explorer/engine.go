package explorer

import (
	bc "github.com/gallactic/hubble_service/blockchain"
	db "github.com/gallactic/hubble_service/database"
)

//Explorer class for connecting block chain to data base
type Explorer struct {
	BCAdapter bc.Adapter
	DBAdapter db.Adapter
}

//Init to initialize database and Sync it with block chain
func (e *Explorer) Init() error {
	//connect to gallactic blockchain by gRPC
	bcAdapter := e.BCAdapter
	clientErr := bcAdapter.CreateGRPCClient()
	if clientErr == nil {
		println("Blockchain client created successfully!")
	} else {
		return clientErr
	}

	//connect to database
	dbAdapter := e.DBAdapter
	connErr := dbAdapter.Connect()
	if connErr != nil {
		return connErr
	}
	println("Connected to database successfully!")
	defer dbAdapter.Disconnect()

	//Sync data with blockchain
	println("Updating blockchain adapter...")
	updateErr := bcAdapter.Update()
	if updateErr != nil {
		return updateErr
	}

	//Get current height of last block
	currentHeight := bcAdapter.GetBlocksLastHeight()

	//Get last block ID that is saved
	lastBlockIDInDB, getLastIDError := dbAdapter.GetBlocksTableLastID()
	if getLastIDError != nil {
		lastBlockIDInDB = 0
	}

	if currentHeight > lastBlockIDInDB {
		println("Reading blocks ", lastBlockIDInDB, " to ", currentHeight, "...")
		//TODO: Read blocks and Save to Data Base
	}

	block, blockErr := bcAdapter.GetBlock(14141)
	if blockErr == nil {
		println("Block Hash :", block.Hash)
	} else {
		println("Block error :", blockErr.Error())
	}

	//println("Last Block -> ", ret3.String())

	//for i := 0; i < 10; i++ {
	//	println("Account ", i, " -> ", len(ret3.)
	//}
	return nil
}
