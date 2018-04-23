package main

// Info about a single ERC20 token
type Token struct {
	// Long name of the token
	Name string `json:"name"`
	// Symbol on exchanges
	Symbol string `json:"symbol"`
	// Address of the token contract
	Address string `json:"address"`
	// Number of decimal points in amounts
	Decimals uint8 `json:"decimals"`
	// Path to logo image
	Logo string `json:"logo"`
}
