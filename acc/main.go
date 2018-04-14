package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatal("Tool must have 2 args")
	}
	fmt.Printf("reading file: %v\n", args[1])
	accs, err := generateAccountsFromRaw(args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reading file: %v\n", args[2])
	balanceMap, err := calcBalancesAfterTransactions(args[2], accs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFinal balances:\n")
	for acc, balance := range balanceMap {
		fmt.Printf("%v,%v\n", acc, balance)
	}
}
