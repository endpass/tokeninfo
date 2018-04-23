package main

import (
	"encoding/json"
	"io"
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
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&tokens); err != nil {
		return nil, err
	}
	results := make([]*Token, len(tokens))
	for i, token := range tokens {
		results[i] = &Token{
			Name:     token.Name,
			Symbol:   token.Symbol,
			Address:  token.Address,
			Decimals: token.Decimals,
		}
	}
	return results, nil
}
