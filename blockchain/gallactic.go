package blockchain

import (
	pb "github.com/gallactic/hubble_service/proto3"
	"google.golang.org/grpc"
)

func GetGRPCGallacticClient() pb.BlockChainClient {
	addr := "68.183.183.19:50500" //TODO: add tConfig.GRPC.ListenAddress
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		println("Blockchain client creation failed!")
		panic(err)
	}
	println("Blockchain client created successfully!")
	return pb.NewBlockChainClient(conn)
}
