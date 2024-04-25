package helper

import (
	"math/rand"
	"strconv"
)

func GenerateGameCode() string {
	return strconv.Itoa(100000 + rand.Intn(9020000))
}

func RandString(l int) string {
	const letterBytes = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`

	b := make([]byte, l)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	
	return string(b)
}