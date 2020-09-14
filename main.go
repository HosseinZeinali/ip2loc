package main

import (
	"github.com/HosseinZeinali/ip2loc/cmd"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Serve()
}
