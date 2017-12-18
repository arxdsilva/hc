package main

import (
	"errors"
	"strconv"
	"strings"
)

type grid struct {
	x, y int64
}

func (g *grid) setGrid(s string) (err error) {
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
	g.x = gridX
	g.y = gridY
	return
}
