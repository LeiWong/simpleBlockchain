package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"
import "encoding/hex"
import "crypto/sha256"
import "encoding/json"

func main() {
	ret := gin.Default()
	ret.GET("/hash", xHash)
	ret.Run(":8888")

}

func xHash(c *gin.Context) {
	hname := c.DefaultQuery("name", "helloworld")
	name := hash(hname)
	value, proof :=proofOfWork(name)
	c.JSON(http.StatusOK, gin.H{"hashName": hname, "powValue": value, "proof":proof})
}

func hash(name string) string {
	h := sha256.New()
	re := make(map[string]interface{})
	re["name"] = name
	// revert map to json and sha256 hash
	jsonname, err := json.Marshal(re)
	if err != nil {
		fmt.Println(err)
	}
	h.Write([]byte(jsonname))
	bs := h.Sum(nil) // Sum accept parameter as salt
	text := hex.EncodeToString(bs)
	fmt.Println(text)
	return text
}

func proofOfWork(hash string) (string, int) {
	proof := 0
	for {
		text, value := validProof(hash, proof)
		if value {
			return text, proof
		} else {
			proof += 1
		}
	}
}

func validProof(hash string, proof int) (string, bool) {
	guess := hash + "|" + string(proof)
	h := sha256.New()
	h.Write([]byte(guess))
	bs := h.Sum(nil)
	text := hex.EncodeToString(bs)
	if text[:4] == "xiyi" {
		return text,true
	} else {
		return text, false
	}
}
