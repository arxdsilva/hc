package main

import (
	"errors"
	"strconv"
	"strings"
)

var ErrAlreadyRegistered = errors.New("d already registered to start in this x and y position")

type grid struct {
	x, y            int64
	droneStartPoint map[int64][]int64
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

func (g *grid) registerD(d dSettings) (err error) {
	_, ok := g.droneStartPoint[d.posX]
	if !ok {
		g.droneStartPoint[d.posX] = append(g.droneStartPoint[d.posX], d.posY)
		return
	}
	for _, v := range g.droneStartPoint[d.posX] {
		if v == d.posY {
			return ErrAlreadyRegistered
		}
	}
	return
}
