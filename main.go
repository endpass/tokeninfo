package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
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

func init() {
	flag.StringVarP(&serverHost, "host", "h", ":8000", "Web server host and port")
}

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
		tokenImages[addr] = filepath.Base(fi.Name())
	}
	return nil
}

// Loads token data from file
func loadTokenList() error {
	f, err := os.Open(tokenListFile)
	if err != nil {
		return err
	}
	tokens, err = ParseTokenList(f)
	if err != nil {
		return err
	}
	// Index tokens and append images
	for _, token := range tokens {
		if tokensBySymbol[token.Symbol] == nil {
			tokensBySymbol[strings.ToUpper(token.Symbol)] = token
		}
		logoImg := tokenImages[strings.ToLower(token.Address)]
		if len(logoImg) > 0 {
			// URL to logo image
			token.Logo = "/img/" + logoImg
		}
	}
	return nil
}

func main() {
	checkErr(readEnv())
	flag.Parse()

	checkErr(loadImageNames())
	checkErr(loadTokenList())
	log.Infof("Cached data for %d tokens", len(tokens))
	log.Infof("Starting server on %s", serverHost)

	startServer()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
