package main

import "fmt"

var Ds = []dSettings{}

func startDsInGrid(g grid) (err error) {
	for {
		d := dSettings{}
		fmt.Printf("\nPlease inform the command sequence for d %v or leave empty to exit:\n", id)
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			break
		}
		err = d.dSettingsFromInput(response)
		if err != nil {
			return
		}
		Ds = append(Ds, d)
	}
	fmt.Println("...exiting\n\n")
	return
}
