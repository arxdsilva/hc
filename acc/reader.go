package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

func generateAccountsFromRaw(file string) (accs map[int]int, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer func() { err = f.Close() }()
	r := csv.NewReader(bufio.NewReader(f))
	rec, err := r.ReadAll()
	if err != nil {
		return
	}
	accs, err = mapAccs(rec)
	if err != nil {
		return
	}
	return
}

func mapAccs(accs [][]string) (map[int]int, error) {
	accsMap := map[int]int{}
	for _, acc := range accs {
		key, err := strconv.Atoi(acc[0])
		if err != nil {
			return nil, err
		}
		value, err := strconv.Atoi(acc[1])
		if err != nil {
			return nil, err
		}
		accsMap[key] = value
	}
	return accsMap, nil
}

func calcBalancesAfterTransactions(file string, accs map[int]int) map[int]int {
	// read transacoes.csv
	// for ea transaction
	// calc balance from acc after transaction
	return accs
}
