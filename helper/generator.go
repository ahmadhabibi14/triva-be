package helper

import (
	"math/rand"
	"strconv"
)

func GenerateGameCode() string {
	return strconv.Itoa(100000 + rand.Intn(9020000))
}