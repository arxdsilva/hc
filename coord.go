package main

import (
	"errors"
	"strconv"
	"strings"
)

const InvalidGrid = errors.New("Invalid Grid")

type dSettings struct {
	grid        []int64
	coords      []int64
	commands    string
	camPosition string
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
	d.coords = append(d.coords, xInt)
	y := s[2:4]
	yInt, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
		return
	}
	d.coords = append(d.coords, yInt)
	d.camPosition = string(s[4])
	d.commands = s[5:]
	return
}

func (d *dSettings) setGrid(s string) (err error) {
	if len(s) == 0 {
		return errors.New("input cannot be empty")
	}
	gridSplitted := strings.Split(s, "x")
	gridX, err := strconv.ParseInt(gridSplitted[0], 10, 64)
	if err != nil {
		return
	}
	gridY, err := strconv.ParseInt(gridSplitted[1], 10, 64)
	if err != nil {
		return
	}
	d.grid = append(d.grid, gridX)
	d.grid = append(d.grid, gridY)
	return
}

func validateGrid(g string) (err error) {
	if len(g) < 5 {
		return errors.New("invalid grid size")
	}
	b := strings.Contains(g, "x")
	if !b {
		return InvalidGrid
	}
	_, err = strconv.ParseInt(g[:2], 10, 64)
	if err != nil {
		return InvalidGrid
	}
	_, err = strconv.ParseInt(g[3:], 10, 64)
	if err != nil {
		return InvalidGrid
	}
	return
}
