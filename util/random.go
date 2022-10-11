package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a util integer [min, max + 1)
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a util string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner(number int) string {
	return RandomString(number)
}

func RandomBalance() int64 {
	return RandomInt(10, 10000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "RIL"}
	return currencies[rand.Intn(len(currencies))]
}
