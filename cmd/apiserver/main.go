package main

import (
	"log"

	"github.com/combos/go-api/internal/app/apiserver"
	"github.com/joho/godotenv"
)

func main() {
	envs, err := godotenv.Read(".env")
	checkError(err)

	config := apiserver.NewConfig(envs)
	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
