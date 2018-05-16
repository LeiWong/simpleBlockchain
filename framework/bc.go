package main

import (
	"blockchain"
	"fmt"
)

func main() {
	var bc blockchain.BlockChain
	block := bc.GenesisBlock()
	b := []blockchain.Block{}
	b = append(b, block)
	bc.Chain = append(bc.Chain, block)
	fmt.Println("ok", b, bc)
}
