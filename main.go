package main

import (
	"context"

	bc "github.com/gallactic/hubble_service/blockchain"
	db "github.com/gallactic/hubble_service/database"
	pb "github.com/gallactic/hubble_service/proto3"
)

//TODO: var tConfig *config.Config
var lastReadBlockNumber int

func main() {

	updateDataBase()

	/*
		ret1, err := client.GetAccounts(context.Background(), &pb.Empty{})
		if err == nil {
			println("Accounts Count -> ", len(ret1.Accounts))
			for i := 0; i < len(ret1.Accounts); i++ {
				println("Account ", i, " -> ", ret1.Accounts[i].Account)
			}
		} else {
			println("GetAccounts Error -> ", err.Error())
		}
	*/

	/*
		ret2, err2 := client.GetGenesis(context.Background(), &pb.Empty{})
		if err2 == nil {
			println("Genesis Hash -> ", ret2.Genesis.ShortHash())
			println("Genesis Chain Name -> ", ret2.Genesis.ChainName())
			println("Genesis Accounts Count -> ", len(ret2.Genesis.Accounts()))
			for i := 0; i < len(ret2.Genesis.Accounts()); i++ {
				println("Account ", i, " -> ", ret2.Genesis.Accounts()[i])
			}

		} else {
			println("GetGenesis Error -> ", err2.Error())
		}
	*/

}

func updateDataBase() error {
	//connect to gallactic blockchain by gRPC
	client := bc.GetGRPCGallacticClient()

	//connect to database
	dbe := db.Postgre{Host: "localhost", Port: 5432, User: "postgres", Password: "gmpinc2007", DBname: "HubbleScan"}
	connErr := dbe.Connect()
	if connErr != nil {
		return connErr
	}
	defer dbe.Disconnect()

	//read blocks from blockchain
	blocks, getBlocksErr := client.GetBlocks(context.Background(), &pb.BlocksRequest{})
	if getBlocksErr != nil {
		return getBlocksErr
	}

	currentHeight := blocks.GetLastHeight()

	lastBlockIDInDB, getLastIDError := dbe.GetBlocksTableLastID()
	if getLastIDError != nil {
		lastBlockIDInDB = 0
	}

	if currentHeight > lastBlockIDInDB {
		println("Reading blocks ", lastBlockIDInDB, " to ", currentHeight, "...")
		//TODO: Read blocks and Save to Data Base
	}

	//println("Last Block -> ", ret3.String())

	//for i := 0; i < 10; i++ {
	//	println("Account ", i, " -> ", len(ret3.)
	//}
	return nil
}
