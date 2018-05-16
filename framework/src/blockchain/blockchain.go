package blockchain

import "time"
import "crypto/sha256"
import "encoding/json"
import "encoding/hex"
import "fmt"

type Transactions struct {
	Sender string
	Recipient string
	Amount float64
}

type Node struct {

}

type Block struct {
	Index int
	Timestamp int64
	Transactions Transactions
	Proof int
	PreviousHash string
}

type BlockChain struct {
	Current_transactions []Transactions
	Chain []Block
	nodes []Node
}

func (chain *BlockChain) NewBlock(proof int, previousHash string) *Block {
	block := Block{}
	block.Index = len(chain.Chain) + 1
	block.Timestamp = time.Now().Unix()
	block.Transactions = Transactions{}
	block.Proof = proof
	block.PreviousHash = previousHash
	chain.Chain = append(chain.Chain, block)
	return &block
}

func (chain BlockChain) ValidBlock() bool {
	lastBlock := chain.Chain[0]
	currentIndex := 1
	for currentIndex < len(chain.Chain) {
		block := chain.Chain[currentIndex]
		if block.PreviousHash != chain.Hash(&lastBlock) {
			return  false
		}
		lastBlock = block
		currentIndex += 1
	}
	return true
}

func (chain *BlockChain) Hash(block *Block) string {
	jsonBytes, err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	h := sha256.New()
	h .Write(jsonBytes)
	bs := h.Sum(nil)
	value := hex.EncodeToString(bs)
	return value
}

func (chain BlockChain) ProofOfWork(block *Block) int {
	lastProof := block.Proof
	lastHash :=chain.Hash(block)
	proof := 0
	for {
		isProof := validProof(lastProof, proof, lastHash)
		if isProof {
			return  proof
		} else {
			proof += 1
		}
	}
}

func validProof(lastProof int, proof int, lasthash string) bool {
	guess := string(lastProof) + "|" + string(proof) + "|" + lasthash
	h := sha256.New()
	h.Write([]byte(guess))
	bs := h.Sum(nil)
	text := hex.EncodeToString(bs)
	if text[:4] == "0000" {
		return true
	} else {
		return false
	}
}

func (chain BlockChain) GenesisBlock() *Block {
	return chain.NewBlock(100, "xiyi")
}
