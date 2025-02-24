package pumpdotfunsdk_test

import (
	"context"
	"os"
	"testing"

	pumpdotfunsdk "github.com/zzispp/pumpdotfun-go-sdk"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

type TestConfig struct {
	rpcClient  *rpc.Client
	wsClient   *ws.Client
	PrivateKey solana.PrivateKey
	mint       solana.PublicKey
}

func GetTestConfig() TestConfig {
	testConfig := TestConfig{}
	testConfig.rpcClient = rpc.New(rpc.DevNet_RPC)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)
	}
	testConfig.wsClient = wsClient
	testConfig.PrivateKey = solana.MustPrivateKeyFromBase58(os.Getenv("TEST_PRIVATE_KEY"))
	testConfig.mint = solana.MustPublicKeyFromBase58(os.Getenv("TEST_MINT"))
	return testConfig
}

func TestBuyToken(t *testing.T) {
	testConfig := GetTestConfig()
	pumpdotfunsdk.SetDevnetMode()
	sig, err := pumpdotfunsdk.BuyToken(
		testConfig.rpcClient,
		testConfig.wsClient,
		testConfig.PrivateKey,
		testConfig.mint,
		10000,
		100,
	)
	if err != nil {
		t.Fatalf("can't buy token: %s", err)
	}
	t.Logf("buy token signature: %s", sig)
}
