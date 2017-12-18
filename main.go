package main

import (
	"fmt"
	"log"
	"os"
)

// mappedGrid represents the x value and y that has been mapped
var mappedGrid = make(map[int64][]int64)

func main() {
	args := os.Args
	if (len(args) == 0) || (len(args) > 2) {
		log.Fatal("bot needs 1 grid")
	}
	err := validateGrid(args[1])
	if err != nil {
		log.Fatal(err)
	}
	g := grid{}
	err = g.setGrid(args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Generating flying grid with dimensions of %vm by %vm\n", g.x, g.y)
	err = startDsInGrid(g)
	if err != nil {
		log.Fatal(err)
	}
}
