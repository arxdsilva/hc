package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidGrid = errors.New("Invalid Grid")
var id = 1

type dSettings struct {
	id            int
	posX          int64
	posY          int64
	commands      string
	camPosition   string
	picturesTaken int
}

func (d *dSettings) dSettingsFromInput(s string) (err error) {
	if len(s) == 0 {
		return errors.New("input cannot be empty")
	}
	coords := s[0:4]
	x := coords[:2]
	xInt, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		return
	}
	d.posX = xInt
	y := s[2:4]
	yInt, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
		return
	}
	d.posY = yInt
	d.camPosition = string(s[4])
	d.commands = s[5:]
	d.id = id
	id++
	d.picturesTaken = 1
	return
}

func validateGrid(g string) (err error) {
	if len(g) < 5 {
		return errors.New("invalid grid size")
	}
	b := strings.Contains(g, "x")
	if !b {
		return ErrInvalidGrid
	}
	_, err = strconv.ParseInt(g[:2], 10, 64)
	if err != nil {
		return ErrInvalidGrid
	}
	_, err = strconv.ParseInt(g[3:], 10, 64)
	if err != nil {
		return ErrInvalidGrid
	}
	return
}

var camPos = map[string]string{
	"N": "Norte",
	"S": "Sul",
	"L": "Leste",
	"O": "Oeste",
}

func (d *dSettings) genReport() {
	fmt.Println("Report:")
	fmt.Printf("- Drone: %v\n", d.id)
	fmt.Printf("  - Final position: [%v,%v]\n", d.posX, d.posY)
	fmt.Printf("  - Direction: %v\n", camPos[d.camPosition])
	fmt.Printf("  - Pictures taken: %v\n\n", d.picturesTaken)
}

func (d *dSettings) rotateCam(r string) {
	switch r {
	case "D":
		if d.camPosition == "N" {
			d.camPosition = "L"
		} else if d.camPosition == "S" {
			d.camPosition = "O"
		} else if d.camPosition == "L" {
			d.camPosition = "S"
		} else if d.camPosition == "O" {
			d.camPosition = "N"
		}
	case "E":
		if d.camPosition == "N" {
			d.camPosition = "O"
		} else if d.camPosition == "S" {
			d.camPosition = "L"
		} else if d.camPosition == "L" {
			d.camPosition = "N"
		} else if d.camPosition == "O" {
			d.camPosition = "S"
		}
	}
}

func (d *dSettings) work(g grid) {
	for _, cmd := range d.commands {
		d.takePhoto(g)
		d.processCmd(string(cmd), g)
	}
}

func (d *dSettings) processCmd(cmd string, g grid) {
	switch cmd {
	case "F":
		d.fly(g)
	default:
		d.rotateCam(cmd)
	}
}

func (d *dSettings) fly(g grid) {
	if (d.camPosition == "N") && ((d.posY + 1) <= g.y) {
		d.posY++
	}
	if (d.camPosition == "S") && ((d.posY - 1) >= 0) {
		d.posY--
	}
	if (d.camPosition == "L") && ((d.posX + 1) <= g.x) {
		d.posX++
	}
	if (d.camPosition == "O") && ((d.posX - 1) >= 0) {
		d.posX--
	}
}

func (d *dSettings) takePhoto(g grid) {
	mapXYPosition(d.posX-1, *d, g)
	mapXYPosition(d.posX, *d, g)
	mapXYPosition(d.posX+1, *d, g)
	d.picturesTaken++
}

func mapXYPosition(pos int64, d dSettings, g grid) {
	if ((pos - 1) > 0) && (pos <= d.posX) {
		_, ok := mappedGrid[pos]
		if !ok {
			if (d.posY - 1) > 0 {
				mappedGrid[pos] = append(mappedGrid[pos], d.posY-1)
			}
			if d.posX != pos {
				mappedGrid[pos] = append(mappedGrid[pos], d.posY)
			}
			if (d.posY + 1) <= (g.y) {
				mappedGrid[pos] = append(mappedGrid[pos], d.posY+1)
			}
		} else {
			var ypos, yM1, yP1 bool
			for _, y := range mappedGrid[pos] {
				if y == d.posY {
					ypos = true
				} else if y == d.posY-1 {
					yM1 = true
				} else if y == d.posY+1 {
					yP1 = true
				}
			}
			if (!ypos) && (pos != d.posX) {
				mappedGrid[pos] = append(mappedGrid[pos], d.posY)
			}
			if !yM1 {
				mappedGrid[pos] = append(mappedGrid[pos], d.posY-1)
			}
			if !yP1 {
				mappedGrid[pos] = append(mappedGrid[pos], d.posY+1)
			}
		}
	}
}
