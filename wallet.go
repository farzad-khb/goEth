package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generatePrivatekey() (*ecdsa.PrivateKey, error) {
	privK, err := crypto.GenerateKey()

	if err != nil {
		return nil, fmt.Errorf("error in generatePrivatekey, can not generate private key: %w", err)
	}
	privkBytes := crypto.FromECDSA(privK)

	fmt.Println("privatekey is:", hexutil.Encode(privkBytes))

	return privK, nil
}
func privateKeyHextoECDSA(privHex string) *ecdsa.PrivateKey {
	priv, err := crypto.HexToECDSA(privHex)
	if err != nil {
		log.Fatalf("error in generateAddressFromPrivHex while decoding hex to privatekey: %w", err)
		return nil

	}
	return priv

}

func generatePublicKey(priv *ecdsa.PrivateKey) *ecdsa.PublicKey {
	pub := priv.Public()
	pubKeyECDSA, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		fmt.Errorf("error in generatePrivatekey, can not assert ecdsa.publickey type : ")
		return nil
	}

	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	fmt.Println("public key is:", hexutil.Encode(pubKeyBytes))

	return pubKeyECDSA

}
func generateAddress(pub *ecdsa.PublicKey) common.Address {
	address := crypto.PubkeyToAddress(*pub)
	fmt.Println("address is:", hexutil.Encode(address[:]))
	return address

}
