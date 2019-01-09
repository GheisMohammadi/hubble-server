package main

import (
	"context"

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
		println("Account -> ", ret1.Accounts[0].Account)
	} else {
		println("GetAccounts Error -> ", err.Error())
	}

	ret2, err2 := client.GetGenesis(context.Background(), &pb.Empty{})
	if err2 == nil {
		println("Genesis -> ", ret2.Genesis)
	} else {
		println("GetGenesis Error -> ", err2.Error())
	}

	ret3, err3 := client.GetBlocks(context.Background(), &pb.BlocksRequest{})
	if err3 == nil {
		println("GetBlocks Size -> ", ret3.Size())
	}
}
