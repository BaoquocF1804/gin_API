package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	//user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    "bao",
		Balance:  100,
		Currency: util.RandomCurrency(),
	}
	var err error
	err = testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	return
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	//for i := 0; i < 10; i++ {
	//	lastAccount = createRandomAccount(t)
	//}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
