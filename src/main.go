package main

import (
	"fmt"
	"net/http"

	"github.com/sawood14012/blockchain-go-node/src/Init"
	"github.com/sawood14012/blockchain-go-node/src/blockchain"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("createBlockChain \tTo Create a new Block Chain")
	fmt.Println("print - Prints the blocks in the chain")
}

func createBlockChain() {
	fmt.Println("\nCreating new BlockChain\n")
	Init.InitializeBlockChain()
	fmt.Println("BlockChain Initialized!")
}

func printChain() {
	iter := blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Student Data: %x\n", block.StudentData)
		fmt.Printf("Signature: %x\n", block.Signature)
		fmt.Printf("Company: %s\n", block.Company)
		fmt.Printf("Verification: %s\n", block.Verification)
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func callcreateBlockChain(w http.ResponseWriter, r *http.Request) {

	createBlockChain()

	w.Header().Set("Content-Type", "application/json")
	message := "BlockChain Initialized!"
	w.Write([]byte(message))
}

func callprintUsage(w http.ResponseWriter, r *http.Request) {

	printUsage()

	w.Header().Set("Content-Type", "application/json")
	message := "Printed Usage!!"
	w.Write([]byte(message))
}
func callprintChain(w http.ResponseWriter, r *http.Request) {

	printChain()

	w.Header().Set("Content-Type", "application/json")
	message := "Printed Chain!!"
	w.Write([]byte(message))
}

func main() {
	port := "8080"
	fmt.Printf("I'm running out of wrods")
	http.HandleFunc("/createBlockChain", callcreateBlockChain)
	http.HandleFunc("/print", callprintChain)
	http.HandleFunc("/usage", callprintUsage)
	fmt.Printf("End points created")
	fmt.Printf("Server listening on localhost:%s\n", port)
	fmt.Printf("Server running")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
