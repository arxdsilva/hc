package main

import (
	"errors"
	"strconv"
)

type dSettings struct {
	gridSize    []int64
	coords      []int64
	commands    string
	camPosition string
}

func dSettingsFromInput(s string) (d dSettings, err error) {
	if len(s) == 0 {
		return d, errors.New("input cannot be empty")
	}
	coords := s[0:4]
	x := coords[:2]
	xInt, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		return d, err
	}
	d.coords = append(d.coords, xInt)
	y := s[2:4]
	yInt, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
		return d, err
	}
	d.coords = append(d.coords, yInt)
	d.camPosition = string(s[4])
	d.commands = s[5:]
	return d, err
}
