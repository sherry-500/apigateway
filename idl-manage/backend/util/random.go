package util

import(
	"strings"
	"time"
	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64{
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomSvcname() string{
	return RandomString(8)
}

func RandomIdl() string {
	return RandomString(10)
}

func RandomType() string{
	types := []string{"PATH", "CONTENT"}
	n := len(types)
	return types[rand.Intn(n)]
}