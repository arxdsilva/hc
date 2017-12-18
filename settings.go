package main

import (
	"errors"
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
