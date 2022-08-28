package pkg

import (
	"math/rand"
	"strconv"
)

func GenerateAccessCode() string {
	accessCode := 999 + rand.Int63n(9000)
	return strconv.Itoa(int(accessCode))
}
