package main

import (
	"fmt"
	"log"
)

// URL
const RINKEBY_URL = "https://rinkeby.infura.io/v3/b4568b4e98344dcba67986115a04834d"

func main() {
	// choose the network
	networkURL := RINKEBY_URL

	// choose privatekey
	var privatekey string
	n := 0
	var err error
	for n != 64 {
		fmt.Println("please enter your privatekey (64-length):")
		_, err = fmt.Scanf("%s", &privatekey)
		n = len(privatekey)
		if err != nil {
			log.Fatalln("can not scan your privatekey", err)
		}

	}

	// choose receiver
	reciever := "0x45C54Af6B8cD6403157B05A1D62a4c7566EB32ad"

	// creating account out of privatekey
	account := generateAddress(generatePublicKey(privateKeyHextoECDSA(privatekey)))

	// balance
	getBalance(networkURL, account)

	// send transaction
	signedtx := signTransaciton(networkURL, privatekey, reciever, 10000000000000000)
	sendTransaction(networkURL, signedtx)

	// get transaction status and amount
	transactionDetial(networkURL, signedtx.Hash().Hex())

}
