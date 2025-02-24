# Pump.fun Go SDK

[![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](#)

A SDK allowing you to create, buy and sell pump.fun tokens in Golang.

## Installation

```bash
go get github.com/zzispp/pumpdotfun-go-sdk
```

## Usage

```go
package main

import (
	"log"
    	"os"
	"context"

	// General solana packages.
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"

	// Pump.fun Go SDK.
	pumpdotfunsdk "github.com/zzispp/pumpdotfun-go-sdk"
)

type Client struct {
	RpcClient *rpc.Client
	WsClient  *ws.Client
}

func getClient(rpcUrl string, wsUrl string) *Client {
	rpcClient := rpc.New(rpcUrl)
	wsClient, err := ws.Connect(context.Background(), wsUrl)
	if err != nil {
		log.Fatalln("ws connection error: ", err)
	}
	return &Client{RpcClient: rpcClient, WsClient: wsClient}
}

func main() {
	privateKey, err := solana.PrivateKeyFromBase58(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalln("please set PRIVATE_KEY environment variable:", privateKey)
	}
	rpcURL := rpc.MainNetBeta_RPC
	wsURL := rpc.MainNetBeta_WS
	c := getClient(rpcURL, wsURL)
	mint := solana.NewWallet()
	_, err = pumpdotfunsdk.CreateToken(
		c.RpcClient,
		c.WsClient,
		privateKey,
		mint,
		"TEST", // symbol
		"TEST", // name
		"https://example.com", // metadata uri
		100000000, // buy 0.1 SOL
		200, // 2% slippage
	)
	if err != nil {
		log.Fatalln("can't create token:", err)
	}
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

BSD-3-Clause license, because FreeBSD is the best OS
