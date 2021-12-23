package ethereumclient

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func newEthereumClient(networkURL string) *ethclient.Client {
	client, err := ethclient.Dial(networkURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection to ", networkURL)
	return client
}

// gets detail of certain transaction by its txid
func GetTransactionDetial(networkURL string, txid string) {
	ec := newEthereumClient(networkURL)
	defer ec.Close()

	tx, pending, err := ec.TransactionByHash(context.Background(), common.HexToHash(txid))
	if err != nil {
		log.Fatalln("can not catch the hash", err)
	}
	if tx == nil {
		log.Fatalln("No transaction found")
	}

	fmt.Println("The transaction amount is ", tx.Value(), " wei")
	fmt.Println("Is the transaction pending:", pending)

}

// gets nonce of an Ethereum account address
func GetNonce(address common.Address, networkURL string) uint64 {
	ec := newEthereumClient(networkURL)
	defer ec.Close()

	nonce, err := ec.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatalln("error in getNonce while calling pending nonce at", err)

		return 0
	}
	return nonce
}

// gets gas price suggestion for the specified network
func GetGasPriceSuggestion(networkURL string) big.Int {
	ec := newEthereumClient(networkURL)
	defer ec.Close()
	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln("error in getNonce while calling pending nonce at", err)

		return big.Int{}
	}
	return *gasPrice

}

// gets chain id of the specified network
func GetChainID(networkURL string) big.Int {
	ec := newEthereumClient(networkURL)
	defer ec.Close()

	chainid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatalln("error in getChainID while getting chain id:", err)
		return big.Int{}
	}
	return *chainid

}

// sends transaction to the network
func SendTransaction(networkURL string, signedtx *types.Transaction) {
	ec := newEthereumClient(networkURL)
	defer ec.Close()

	if err := ec.SendTransaction(context.Background(), signedtx); err != nil {
		log.Fatalln("error in getChainID while getting chain id:", err)

	}
	fmt.Println("transaction sent. txid: ", signedtx.Hash().Hex())

}

// gets balance of a certain account
func GetBalance(networkURL string, account common.Address) {
	ec := newEthereumClient(networkURL)
	defer ec.Close()

	balance, err := ec.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalln("error in getBalance", err)
	}
	fmt.Println("for this account, the balance is: ", balance)

}
