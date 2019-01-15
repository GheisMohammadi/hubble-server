package tests

import (
	"testing"

	db "github.com/gallactic/hubble_service/database"
	"github.com/stretchr/testify/require"
)

func TestAccountsInDataBase(t *testing.T) {

	dbe := db.Postgre{Host: "localhost", Port: 5432, User: "postgres", Password: "gmpinc2007", DBname: "HubbleScan"}
	connErr := dbe.Connect()
	require.NoError(t, connErr)

	defer dbe.Disconnect()

	acc := db.Account{Address: "Addr123456", PublicKey: "ABC", Balance: 1234.56, Permission: "Perm456", Sequence: 2, Code: "CodeF1F2"}
	insertErr := dbe.InsertAccount(&acc)
	require.NoError(t, insertErr)

	sAcc, GAccErr := dbe.GetAccount(7)
	require.NoError(t, GAccErr)
	require.Equal(t, sAcc.ID, 7)

	sAcc, GNoAccErr := dbe.GetAccount(10000000)
	require.Error(t, GNoAccErr)
}
