package blockchain

import (
	"context"
	"fmt"

	pb "github.com/gallactic/hubble_server/proto3"
	"google.golang.org/grpc"
)

//Gallactic class for connecting to Gallactic block chain
type Gallactic struct {
	client   *pb.BlockChainClient
	blocks   *pb.BlocksResponse
	accounts *pb.AccountsResponse
}

//CreateGRPCClient creates a client for communicating with gallactic blockchain
func (g *Gallactic) CreateGRPCClient() error {
	addr := "68.183.183.19:50500" //TODO: add tConfig.GRPC.ListenAddress
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewBlockChainClient(conn)
	g.client = &client
	return nil
}

//Update will refresh all data and sync with block chain
func (g *Gallactic) Update() error {
	var getBlocksErr error
	var getAccountsErr error
	client := *g.client

	g.blocks, getBlocksErr = client.GetBlocks(context.Background(), &pb.BlocksRequest{})
	if getBlocksErr != nil {
		return getBlocksErr
	}

	g.accounts, getAccountsErr = client.GetAccounts(context.Background(), &pb.Empty{})
	if getAccountsErr != nil {
		return getAccountsErr
	}
	return nil
}

//GetBlocksLastHeight return last height
func (g *Gallactic) GetBlocksLastHeight() uint64 {
	return g.blocks.GetLastHeight()
}

//GetBlockMeta returns specified block
func (g *Gallactic) GetBlockMeta(height uint64) (*BlockMeta, error) {

	client := *g.client
	blockRes, getBlockErr := client.GetBlock(context.Background(), &pb.BlockRequest{Height: height})
	if getBlockErr != nil {
		return nil, getBlockErr
	}

	blockMeta := blockRes.BlockMeta

	var meta BlockMeta
	toBlockMeta(blockMeta, &meta)
	return &meta, nil
}

//GetBlock returns specified block
func (g *Gallactic) GetBlock(height uint64) (*Block, error) {

	if height > g.blocks.GetLastHeight() {
		return nil, fmt.Errorf("block height out of range (max is " + string(g.blocks.GetLastHeight()) + ")")
	}

	client := *g.client
	blockRes, getBlockErr := client.GetBlock(context.Background(), &pb.BlockRequest{Height: height})
	if getBlockErr != nil {
		return nil, getBlockErr
	}

	var b Block
	toBlock(blockRes, &b)
	return &b, nil
}

//GetAccountsCount returns number of accounts
func (g *Gallactic) GetAccountsCount() int {
	l := len(g.accounts.Accounts)
	return l
}

//GetAccount returns specified account
func (g *Gallactic) GetAccount(id int) (*Account, error) {
	acc := g.accounts.Accounts[id].Account
	ID := uint64(id)
	var retAcc Account
	toAccount(acc, &retAcc)
	retAcc.ID = ID
	return &retAcc, nil
}

//GetAccounts returns all accounts in array of accounts
func (g *Gallactic) GetAccounts() ([]*Account, error) {
	l := len(g.accounts.Accounts)

	retAccounts := make([]*Account, l)

	for i := 0; i < l; i++ {
		acc := g.accounts.Accounts[i].Account
		ID := uint64(i)
		toAccount(acc, retAccounts[i])
		retAccounts[i].ID = ID
	}

	return retAccounts, nil
}

//GetBlocksMeta returns a group of blocks for faster access them
func (g *Gallactic) GetBlocksMeta(from uint64, to uint64) ([]*BlockMeta, error) {
	client := *g.client
	blocks, getBlocksErr := client.GetBlocks(context.Background(), &pb.BlocksRequest{MinHeight: from, MaxHeight: to})
	if getBlocksErr != nil {
		return nil, getBlocksErr
	}

	n := len(blocks.BlockMeta)
	retBlocks := make([]*BlockMeta, 0, n)
	for i := 0; i < n; i++ {
		toBlockMeta(&blocks.BlockMeta[i], retBlocks[i])
	}

	return retBlocks, nil
}

//GetBlocks returns a group of blocks for faster access them
func (g *Gallactic) GetBlocks(from uint64, to uint64) ([]*Block, error) {
	client := *g.client
	blocks, getBlocksErr := client.GetBlocks(context.Background(), &pb.BlocksRequest{MinHeight: from, MaxHeight: to})
	if getBlocksErr != nil {
		return nil, getBlocksErr
	}

	n := len(blocks.BlockMeta)
	retBlocks := make([]*Block, n)

	for i := 0; i < n; i++ {
		retBlocks[i] = new(Block)
		BlockMetaToBlock(&blocks.BlockMeta[i], retBlocks[i])
	}

	return retBlocks, nil
}
