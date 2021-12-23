package main

import (
	"eth/config"
	"eth/ethereumclient"
	"fmt"
	"log"
)

func main() {
	// choose the network
	networkURL := config.GetEnv("RINKEBY_HTTP_URL")
	fmt.Println("network URL from config is", networkURL)

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
	ethereumclient.GetBalance(networkURL, account)

	// send transaction
	signedtx := signTransaciton(networkURL, privatekey, reciever, 10000000000000000)
	ethereumclient.SendTransaction(networkURL, signedtx)

	// get transaction status and amount
	ethereumclient.GetTransactionDetial(networkURL, signedtx.Hash().Hex())

}
