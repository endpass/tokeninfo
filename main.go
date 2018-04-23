package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	// Path to JSON file containing list of tokens
	tokenListFile string
	// Directory where token images are stored
	tokenImageDir string

	// Map of token addresses to paths to token images
	tokenImages = map[string]string{}
)

// Read config from environment variables
func readEnv() error {
	tokenListFile = os.Getenv("TOKEN_LIST")
	if tokenListFile == "" {
		return errors.New("$TOKEN_LIST must not be empty")
	}
	tokenImageDir = os.Getenv("TOKEN_IMAGE_DIR")
	if tokenImageDir == "" {
		return errors.New("$TOKEN_IMAGE_DIR must not be empty")
	}
	return nil
}

func main() {
	if err := readEnv(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("token-list")
}
