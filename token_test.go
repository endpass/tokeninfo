package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testList = `
[
	{
		"symbol": "GIF",
		"address": "0xFcD862985628b254061F7A918035B80340D045d3",
		"decimals": 18,
		"name": "GIFcoin Token",
		"ens_address": "",
		"website": "https://gifcoin.io/",
		"logo": {
			"src":
				"https://www.gifcoin.io/assets/images/default/frontend/gifsingle400.png",
			"width": "400",
			"height": "400",
			"ipfs_hash": ""
		},
		"support": {
			"email": "support@gifcoin.io",
			"url": ""
		},
		"social": {
			"blog": "",
			"chat": "",
			"facebook": "https://facebook.com/gifcoin.io",
			"forum": "",
			"github": "",
			"gitter": "",
			"instagram": "",
			"linkedin": "",
			"reddit": "",
			"slack": "",
			"telegram": "https://t.me/gifcoin",
			"twitter": "https://twitter.com/gifcoin_io",
			"youtube": "https://youtube.com/channel/UCLq13wzOH1STqW8I-Z-ctAQ"
		}
	},
	{
		"symbol": "DTT",
		"address": "0xf9F7c29CFdf19FCf1f2AA6B84aA367Bcf1bD1676",
		"decimals": 18,
		"name": "Delphi Tech Token",
		"ens_address": "",
		"website": "https://delphifund.org/",
		"logo": {
			"src":
				"https://delphifund.org/wp-content/uploads/2018/04/delphi400x400.jpg",
			"width": "400",
			"height": "400",
			"ipfs_hash": ""
		},
		"support": {
			"email": "support@delphifund.org",
			"url": ""
		},
		"social": {
			"blog": "",
			"chat": "",
			"facebook": "",
			"forum": "",
			"github": "https://github.com/DTToken",
			"gitter": "",
			"instagram": "",
			"linkedin": "",
			"reddit": "",
			"slack": "",
			"telegram": "",
			"twitter": "https://twitter.com/Delphitechtoken",
			"youtube": ""
		}
	},
	{
		"symbol": "DTT2",
		"address": "0xf9f7c29cfdf19fcf1f2aa6b84aa367bcf1bd1676",
		"decimals": 18,
		"name": "Delphi Tech Token 2"
	}
]
`

func TestTokenList(t *testing.T) {
	assert := assert.New(t)
	r := strings.NewReader(testList)
	tokens, err := ParseTokenList(r)
	assert.NoError(err)
	assert.Len(tokens, 2)
	assert.Equal("GIFcoin Token", tokens[0].Name)
	assert.Equal("0xFcD862985628b254061F7A918035B80340D045d3", tokens[0].Address)
	assert.Equal("GIF", tokens[0].Symbol)
	assert.Equal(18, tokens[0].Decimals)
	assert.Equal(tokens[1].Symbol, "DTT")
}
