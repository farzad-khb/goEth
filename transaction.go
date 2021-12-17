package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func signTransaciton(networkURL string, privHex string, to string, value int64) *types.Transaction {

	// making fromAddress from private key
	fromAddress := generateAddress(generatePublicKey(privateKeyHextoECDSA(privHex)))

	// getting nonce from network
	nonce := getNonce(fromAddress, networkURL)

	// value
	amount := big.NewInt(value)

	// gas limit is set 21000 units
	gasLimit := uint64(21000)

	// gas price
	gasPrice := getGasPriceSuggestion(networkURL)

	// address
	toAddress := common.HexToAddress(to)

	// getting chain id
	chainID := getChainID(networkURL)

	// making transaction
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, &gasPrice, []byte{})

	// signing transaction
	signTX, err := types.SignTx(tx, types.NewEIP155Signer(&chainID), privateKeyHextoECDSA(privHex))
	if err != nil {
		log.Fatal("error in signTransaciton while signTX: ", err)

		return nil
	}
	fmt.Println("transaction successfully signed")
	return signTX
}
