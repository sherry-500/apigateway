package db

import (
	"context"
	"database/sql"
	"idlManage/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomIdlmap(t *testing.T) Idlmap{
	arg := CreateIdlmapParams{
		Svcname: util.RandomSvcname(),
		Idl: util.RandomIdl(),
		Type: util.RandomType(),
	}

	idlmap, err := testQueries.CreateIdlmap(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, idlmap)

	require.Equal(t, arg.Svcname, idlmap.Svcname)
	require.Equal(t, arg.Idl, idlmap.Idl)
	require.Equal(t, arg.Type, idlmap.Type)

	require.NotZero(t, idlmap.ID)
	require.NotZero(t, idlmap.CreatedAt)

	return idlmap
}

func TestCreateIdlmap(t *testing.T){
	createRandomIdlmap(t)
}

func TestGetIdlmap(t *testing.T){
	idlmap1 := createRandomIdlmap(t)
	idlmap2, err := testQueries.GetIdlmap(context.Background(), idlmap1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, idlmap2)

	require.Equal(t, idlmap1.ID, idlmap2.ID)
	require.Equal(t, idlmap1.Svcname, idlmap2.Svcname)
	require.Equal(t, idlmap1.Idl, idlmap2.Idl)
	require.WithinDuration(t, idlmap1.CreatedAt, idlmap2.CreatedAt, time.Second)
}

func TestGetIdlmapByName(t *testing.T){
	idlmap1 := createRandomIdlmap(t)
	idlmap2, err := testQueries.GetIdlmapByName(context.Background(), idlmap1.Svcname)
	require.NoError(t, err)
	require.NotEmpty(t, idlmap2)

	require.Equal(t, idlmap1.ID, idlmap2.ID)
	require.Equal(t, idlmap1.Svcname, idlmap2.Svcname)
	require.Equal(t, idlmap1.Idl, idlmap2.Idl)
	require.WithinDuration(t, idlmap1.CreatedAt, idlmap2.CreatedAt, time.Second)
}

func TestUpdateIdlmap(t *testing.T){
	idlmap1 := createRandomIdlmap(t)

	arg := UpdateIdlmapParams{
		Svcname: idlmap1.Svcname,
		Idl: idlmap1.Idl,
		Type: idlmap1.Type,
	}

	idlmap2, err := testQueries.UpdateIdlmap(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, idlmap2)

	require.Equal(t, idlmap1.ID, idlmap2.ID)
	require.Equal(t, idlmap1.Svcname, idlmap2.Svcname)
	require.Equal(t, arg.Idl, idlmap2.Idl)
	require.Equal(t, idlmap1.Type, idlmap2.Type)
	require.WithinDuration(t, idlmap1.CreatedAt, idlmap2.CreatedAt, time.Second)
}

func TestDeleteIdlmap(t *testing.T){
	idlmap1 := createRandomIdlmap(t)
	err := testQueries.DeleteIdlmap(context.Background(), idlmap1.Svcname)
	require.NoError(t, err)

	idlmap2, err := testQueries.GetIdlmap(context.Background(), idlmap1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, idlmap2)
}

func TestListIdlmaps(t *testing.T){
	for i := 0; i < 10; i++{
		createRandomIdlmap(t)
	}

	arg := ListIdlmapParams{
		Offset: 5,
		Limit: 5,
	}
	
	idlmaps, err := testQueries.ListIdlmap(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, idlmaps, 5)

	for _, idlmap:= range idlmaps{
		require.NotEmpty(t, idlmap)
	}
}


