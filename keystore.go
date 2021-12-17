package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func generateKeyStore(keydir string) {
	ks := keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "1234"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatalln("error is generateKeyStore while generating new account", err)
	}
	fmt.Println("this is the keystore account:", account.Address.Hex())
}
func importKeyStore(file string) {
	ks := keystore.NewKeyStore("./temp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "1234"
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("error is importKeyStore while reading file: ", err)
	}
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatalln("error is importKeyStore while creating account: ", err)
	}
	fmt.Println("this is the keystore account:", account.Address.Hex())
	err = os.Remove(file)
	if err != nil {
		log.Fatalln("error is importKeyStore while removing file", err)
	}
}
