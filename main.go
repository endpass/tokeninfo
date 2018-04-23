package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Config variables
var (
	// Path to JSON file containing list of tokens
	tokenListFile string
	// Directory where token images are stored
	tokenImageDir string
	// Host for server
	serverHost string
)

var (
	// Map of token addresses to paths to token images
	tokenImages = map[string]string{}

	// token data loaded from file, indexed by symbol
	tokens         = []*Token{}
	tokensBySymbol = map[string]*Token{}
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

// Loads normalized paths to images into memory
func loadImageNames() error {
	fis, err := ioutil.ReadDir(tokenImageDir)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		if !fi.Mode().IsRegular() {
			continue
		}
		addr := strings.ToLower(strings.TrimSuffix(fi.Name(), filepath.Ext(fi.Name())))
		tokenImages[addr] = fi.Name()
	}
	return nil
}

func main() {
	checkErr(readEnv())
	checkErr(loadImageNames())
	log.Infof("Starting server on %s", serverHost)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
