package main

import (
	"errors"
	"strconv"
	"strings"
)

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
