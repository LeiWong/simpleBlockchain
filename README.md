# Simple Blockchain implement by golang

## download latest version golang
> download go1.9.2.linux.tar.gz
> cd /home/ubuntu
> tar -zxvf go1.9.2.linux-386.tar.gz

## setting environment variable
> export GOROOT=/home/ubuntu/go
> export PATH=$GOROOT/bin:$PATH
> export GOPATH=/home/ubuntu/go/gopkg
> go version

## install dependent package
> go get github.com/gin-gonic/gin

## run instance
> go run main.go

## mining blockachain
> curl 0.0.0.0:9090/v1/mine?name=golang
```
return json data

{
    "block":{
                "Index":3,
                "Timestamp":1526349942,
                "Transactions":{"Sender":"","Recipient":"","Amount":0},
                "Proof":16065,
                "PreviousHash":"34b8357f0a2b131fdb22f47b0a40230eb8d2c0a46625b7c5007be6addca6d29e"
                },
    "message":"New block"
    }
```
