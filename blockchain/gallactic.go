package blockchain

import (
	"context"
	"encoding/hex"
	"fmt"

	pb "github.com/gallactic/hubble_service/proto3"
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

//GetBlocksMeta returns specified block
func (g *Gallactic) GetBlocksMeta(id int) (*BlockMeta, error) {

	if id >= len(g.blocks.BlockMeta) {
		return nil, fmt.Errorf("block index out of range (max is " + string(len(g.blocks.BlockMeta)) + ")")
	}

	meta := &g.blocks.BlockMeta[id]

	return toBlockMeta(meta), nil
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

	blockMeta := blockRes.BlockMeta
	blockID := blockMeta.BlockID
	header := blockRes.Block.Header
	data := blockRes.Block.Data

	var b *Block
	b = &Block{
		ID:            height,
		Hash:          hex.EncodeToString(blockID.Hash),
		ChainID:       header.ChainID,
		Height:        header.Height,
		Time:          header.Time.String(),
		LastBlockHash: hex.EncodeToString(header.LastBlockID.Hash),
		TxCounts:      header.NumTxs,
		DataHash:      hex.EncodeToString(data.Hash())}

	return b, nil
}

//GetAccountsCount returns number of accounts
func (g *Gallactic) GetAccountsCount() int {
	l := len(g.accounts.Accounts)
	return l
}

//GetAccount returns specified account
func (g *Gallactic) GetAccount(id int) (*Account, error) {
	acc := g.accounts.Accounts[id].Account
	code := fmt.Sprintf("%s", acc.Code())
	ID := uint64(id)
	hsAcc := &Account{Address: acc.Address().String(), PublicKey: "", Balance: acc.Balance(),
		Permission: acc.Permissions().String(), Sequence: acc.Sequence(), Code: code, ID: ID}
	return hsAcc, nil
}

/*
for i := 0; i < len(ret1.Accounts); i++ {
	println("Account ", i, " -> ", ret1.Accounts[i].Account)
}
*/
