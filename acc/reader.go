package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

func generateAccountsFromRaw(file string) (accs map[int64]int64, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	r := csv.NewReader(bufio.NewReader(f))
	_, err = r.ReadAll()
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
