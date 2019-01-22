package tests

import (
	"testing"

	bc "github.com/gallactic/hubble_server/blockchain"
	"github.com/stretchr/testify/require"
)

func TestBlockChain(t *testing.T) {

	clientErr := bc.CreateGRPCClient()
	require.NoError(t, clientErr)

	updateErr := bc.Update()
	require.NoError(t, updateErr)

	/*
		ret1, err := bc.GetAccounts()
		require.NoError(t, err)

		numAccounts := len(ret1.Accounts)
		require.NotEqual(t, numAccounts, 0)

		_, err2 := client.GetGenesis(context.Background(), &pb.Empty{})
		require.NoError(t, err2)

		_, err3 := client.GetBlocks(context.Background(), &pb.BlocksRequest{})
		require.NoError(t, err3)
	*/
}
