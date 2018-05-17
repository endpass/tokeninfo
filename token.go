package main

import (
	"encoding/json"
	"io"
	"strings"
)

// Info about a single ERC20 token
type Token struct {
	// Long name of the token
	Name string `json:"name"`
	// Symbol on exchanges
	Symbol string `json:"symbol"`
	// Address of the token contract
	Address string `json:"address"`
	// Number of decimal points in amounts
	Decimals int `json:"decimals"`
	// Path to logo image
	Logo string `json:"logo"`
}

// Token info from
// https://github.com/MyEtherWallet/ethereum-lists/blob/master/tokens/tokens-eth.json
type tokenListToken struct {
	// Long name of the token
	Name string `json:"name"`
	// Symbol on exchanges
	Symbol string `json:"symbol"`
	// Address of the token contract
	Address string `json:"address"`
	// Number of decimal points in amounts
	Decimals int `json:"decimals"`
	// Path to logo image
	Logo string `json:"-"`
}

func ParseTokenList(r io.Reader) ([]*Token, error) {
	var tokens []*tokenListToken
	// To enforce uniqueness by address, only take the first one at an
	// address
	tokenAddrs := make(map[string]bool)

	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&tokens); err != nil {
		return nil, err
	}
	results := make([]*Token, 0)
	for _, token := range tokens {
		// Skip if not unique
		addr := strings.ToLower(token.Address)
		if tokenAddrs[addr] {
			continue
		}
		tokenAddrs[addr] = true
		result := &Token{
			Name:     token.Name,
			Symbol:   token.Symbol,
			Address:  token.Address,
			Decimals: token.Decimals,
		}
		results = append(results, result)
	}
	return results, nil
}
