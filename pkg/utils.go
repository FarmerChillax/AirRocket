package pkg

import (
	"log"
	"math/rand"
	"os"
	"strconv"
)

func GenerateAccessCode() string {
	accessCode := 999 + rand.Int63n(9000)
	return strconv.Itoa(int(accessCode))
}

func CheckFileExist(URI string) bool {
	fileInfo, err := os.Stat(URI)
	if err != nil {
		// if errors.Is(err, os.ErrNotExist) {
		// 	return false
		// }
		return false
	}

	if fileInfo.IsDir() {
		log.Printf("URI is a directory")
		return false
	}
	return true
}
