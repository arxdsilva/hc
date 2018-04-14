package main

import (
	"bufio"
	"encoding/csv"
	"io"
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

func calcBalancesAfterTransactions(file string, accs map[int]int) (map[int]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() { err = f.Close() }()
	accs, err = calculateBalance(accs, csv.NewReader(bufio.NewReader(f)))
	if err != nil {
		return nil, err
	}
	return accs, err
}

func calculateBalance(accs map[int]int, r *csv.Reader) (m map[int]int, err error) {
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		accID, err := strconv.Atoi(rec[0])
		if err != nil {
			return nil, err
		}
		transaction, err := strconv.Atoi(rec[1])
		if err != nil {
			return nil, err
		}
		if _, ok := accs[accID]; !ok {
			continue
		}
		accs[accID] = accs[accID] + transaction
	}
	m = accs
	return
}
