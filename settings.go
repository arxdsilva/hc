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

func (d *dSettings) work() {
	for {

	}
}
