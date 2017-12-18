package main

import (
	"log"
	"os"
)

var mappedGrid map[int64][]int64

func main() {
	args := os.Args
	if (len(args) == 0) || (len(args) > 1) {
		log.Fatal("bot needs 1 grid")
	}
	err := validateGrid(args[0])
	if err != nil {
		log.Fatal(err)
	}
	d := dSettings{}
	err = d.setGrid(args[0])
	if err != nil {
		log.Fatal(err)
	}

}
