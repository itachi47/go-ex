package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system env")
	}

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
}
