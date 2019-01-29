package main

import (
	"time"

	bc "github.com/gallactic/hubble_server/blockchain"
	db "github.com/gallactic/hubble_server/database"
	ex "github.com/gallactic/hubble_server/explorer"
)

//TODO: var tConfig *config.Config
var lastReadBlockNumber int

func main() {

	start()

}

func start() {
	bcAdapter := bc.Gallactic{}
	dbAdapter := db.Postgre{Host: "localhost", Port: 5432, User: "postgres", Password: "123456", DBname: "HubbleScan"}
	explorerEngine := ex.Explorer{BCAdapter: &bcAdapter, DBAdapter: &dbAdapter}

	explorerEngine.Init()

	for {

		explorerEngine.Update()
		time.Sleep(1000 * time.Millisecond)

	}

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
