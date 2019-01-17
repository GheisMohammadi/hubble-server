package main

import (
	bc "github.com/gallactic/hubble_service/blockchain"
	db "github.com/gallactic/hubble_service/database"
	ex "github.com/gallactic/hubble_service/explorer"
)

//TODO: var tConfig *config.Config
var lastReadBlockNumber int

func main() {

	bcAdapter := bc.Gallactic{}
	dbAdapter := db.Postgre{Host: "localhost", Port: 5432, User: "postgres", Password: "gmpinc2007", DBname: "HubbleScan"}
	explorerEngine := ex.Explorer{BCAdapter: &bcAdapter, DBAdapter: &dbAdapter}

	explorerEngine.Init()

	/*

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
