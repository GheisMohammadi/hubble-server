package tests

import (
	"context"
	"testing"

	bc "github.com/gallactic/hubble_service/blockchain"
	pb "github.com/gallactic/hubble_service/proto3"
	"github.com/stretchr/testify/require"
)

func TestBlockChain(t *testing.T) {

	client := bc.GetGRPCGallacticClient()

	ret1, err := client.GetAccounts(context.Background(), &pb.Empty{})
	require.NoError(t, err)

	numAccounts := len(ret1.Accounts)
	require.NotEqual(t, numAccounts, 0)

	_, err2 := client.GetGenesis(context.Background(), &pb.Empty{})
	require.NoError(t, err2)

	_, err3 := client.GetBlocks(context.Background(), &pb.BlocksRequest{})
	require.NoError(t, err3)
}
