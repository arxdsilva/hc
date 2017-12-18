package main

import "fmt"

func startDsInGrid(g grid) (err error) {
	for {
		d := dSettings{}
		fmt.Printf("\nPlease inform the command sequence for d %v or leave empty to exit:\n", id)
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			return
		}
		d.dSettingsFromInput(response)
	}
	return
}
