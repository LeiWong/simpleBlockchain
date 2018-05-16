package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"net/http"
	"blockchain"
	"fmt"
)

func main() {
	bc := blockchain.BlockChain{}
	block := bc.GenesisBlock()
	fmt.Println(bc.Chain)
	bc.Chain = append(bc.Chain, *block)
	// write log file
	gin.DisableConsoleColor()
	// Logging to a file
	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()

	// Simple group :v1
	v1 := router.Group("/v1")
	{
		v1.GET("/mine", func (c *gin.Context) {
		l := len(bc.Chain)
		lastBlock := bc.Chain[l-1]
		proof := bc.ProofOfWork(&lastBlock)
		trans := blockchain.Transactions{"hh", "xiyi", 100}
		bc.Current_transactions = append(bc.Current_transactions, trans)
		previousHash := bc.Hash(&lastBlock)
		var block *blockchain.Block
		block = bc.NewBlock(proof, previousHash)
		c.JSON(http.StatusOK, gin.H{
			"block": block, 
			"message": "New block"})
	})

	v1.GET("/chain", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"chain": bc.Chain, "message": "ok"})
 })
 }

// 	 run server
         router.Run(":9090")
}
