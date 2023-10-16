package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Read from a file and return an object with all the configs.
func Read() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf(".env file not readed '%s'\n", err)
		fmt.Println("It is expected that the environment variables are already loaded (normal case with docker)")
	}

	return nil
}
