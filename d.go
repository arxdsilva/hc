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
		err = g.registerD(d)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Ds = append(Ds, d)
		fmt.Printf("- D initialized at position [%v, %v] facing \"%v\", with sequence \"%v\" \n", d.posX, d.posY, camPos[d.camPosition], d.commands)
	}
	fmt.Println("...exiting\n\n")
	fmt.Println("Report:")
	for _, d := range Ds {
		d.work(g)
		d.genReport()
	}
	return
}
