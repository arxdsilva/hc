package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var ErrInvalidGrid = errors.New("Invalid Grid")

type dSettings struct {
	posX        int64
	posY        int64
	gridMaxX    int64
	gridMaxY    int64
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
	d.posX = xInt
	y := s[2:4]
	yInt, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
		return
	}
	d.posY = yInt
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
	d.gridMaxX = gridX
	d.gridMaxY = gridY
	return
}

func (d *dSettings) calcMinDFromSettings() int64 {
	var smaller int64
	if d.gridMaxX > d.gridMaxY {
		smaller = d.gridMaxY
	} else {
		smaller = d.gridMaxX
	}
	minDNum := math.Abs(float64(smaller / 3.0))
	remainder := smaller % 3.0
	if remainder >= 1 {
		minDNum++
	}
	return int64(minDNum)
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