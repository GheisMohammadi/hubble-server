package main

import (
	"context"

	db "github.com/gallactic/hubble_service/database"
	pb "github.com/gallactic/hubble_service/proto3"

	"google.golang.org/grpc"
)

//TODO: var tConfig *config.Config

func grpcBlockchainClient() pb.BlockChainClient {
	addr := "68.183.183.19:50500" //TODO: add tConfig.GRPC.ListenAddress
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		println("Blockchain client creation failed!")
		panic(err)
	}
	println("Blockchain client created successfully!")
	return pb.NewBlockChainClient(conn)
}

func main() {
	client := grpcBlockchainClient()

	ret1, err := client.GetAccounts(context.Background(), &pb.Empty{})
	if err == nil {
		println("Accounts Count -> ", len(ret1.Accounts))
		for i := 0; i < len(ret1.Accounts); i++ {
			println("Account ", i, " -> ", ret1.Accounts[i].Account)
		}
	} else {
		println("GetAccounts Error -> ", err.Error())
	}

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

	ret3, err3 := client.GetBlocks(context.Background(), &pb.BlocksRequest{})
	if err3 == nil {
		println("GetBlocks Size -> ", ret3.GetLastHeight())

		//println("Last Block -> ", ret3.String())

		//for i := 0; i < 10; i++ {
		//	println("Account ", i, " -> ", len(ret3.)
		//}

	}

	dbe := db.DBPostgre{Host: "localhost", Port: 5432, User: "postgres", Password: "gmpinc2007", DBname: "HubbleScan"}
	connErr := dbe.Connect()
	if connErr == nil {
		println("Connected Successfully!")
	} else {
		println("connection error: ", connErr)
	}
	defer dbe.Disconnect()

	acc := db.Account{Address: "Addr123456", PublicKey: "ABC", Balance: 1234.56, Permission: "Perm456", Sequence: 2, Code: "CodeF1F2"}
	dbe.InsertAccount(&acc)

	sAcc, GAccErr := dbe.GetAccount(7)
	if GAccErr == nil {
		println("Account: ", sAcc.Address)
	} else {
		println("Get Account error: ", GAccErr.Error())
	}
}
