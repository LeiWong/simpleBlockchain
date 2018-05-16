package handler

import (
	"blockchain"
)

bc := blockchain.BlockChain{}
bc.Genesis()

func Mine(c *gin.Context) {
	lastBlock := bc.chain[-1]
	proof := bc.ProofOfWork(lastBlock)
	bc.Transaction = blockchain.Transaction{
		"hh", "kk", 0}
	previousHash := handler.Hash(lastBlock)
	block := bc.NewBlock(proof, previousHash)
	
	c.JSON(http.StatusOK, gin.H{"block": block, "message": "New block foged"})

}
