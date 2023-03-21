package main

import (
	"log"

	"github.com/FarmerChillax/AirRocket/startup"
)

const ADDRESS = "127.0.0.1:5000"

func main() {
	err := startup.StartUp()
	if err != nil {
		log.Printf("startup.StartUp err: %v", err)
	}
}
