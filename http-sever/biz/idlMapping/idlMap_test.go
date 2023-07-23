package idlMapping

import (
	"math/rand"
	"strings"
	"testing"
	"os"

	"github.com/stretchr/testify/require"
)

//these tests cannot test InitMap
//they just test the crud operation on the InitMap, but cannot test the crud operation on the idlPath.txt
//you have to check the idlPath.txt's content after running these tests

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestMain(m *testing.M) {
	InitMap()
	os.Exit(m.Run())
}

func TestAddIdl(t *testing.T) {
	svcName := randomString(12)
	idlPath1 := randomString(10)

	result := AddIdl(svcName, idlPath1)
	idlPath2, ok := IDLMap[svcName]

	require.Equal(t, true, result)
	require.Equal(t, true, ok)
	require.NotEmpty(t, idlPath2)
	require.Equal(t, idlPath1, idlPath2)
}

func TestGetIdl(t *testing.T){
	svcName1 := randomString(12)
	idlPath1 := randomString(10)

	AddIdl(svcName1, idlPath1)
	idlPath2 := GetIdl(svcName1)

	require.NotEmpty(t, idlPath2)
	require.Equal(t, idlPath1, idlPath2)
}

func TestUpdateIdl(t *testing.T){
	svcName1 := randomString(12)
	idlPath1 := randomString(10)
	AddIdl(svcName1, idlPath1)

	idlPath2 := randomString(14)
	UpdateIdl(svcName1, idlPath2)
	idlPath3 := GetIdl(svcName1)

	require.NotEmpty(t, idlPath3)
	require.Equal(t, idlPath2, idlPath3)
}

func TestDelIdl(t *testing.T) {
	svcName1 := randomString(12)
	idlPath1 := randomString(10)
	AddIdl(svcName1, idlPath1)

	DelIdl(svcName1)
	idlPath2 := GetIdl(svcName1)

	require.Empty(t, idlPath2)
}


