package main

import (
	"fmt"

	check "gopkg.in/check.v1"
)

func (s *S) TestStripInput(c *check.C) {
	dSett := dSettings{}
	err := dSett.dSettingsFromInput("")
	c.Assert(err, check.NotNil)
	err = dSett.dSettingsFromInput("0315ODEDEFFDFFD")
	c.Assert(dSett.posX, check.NotNil)
	c.Assert(dSett.posY, check.NotNil)
	c.Assert(dSett.camPosition, check.NotNil)
	c.Assert(dSett.commands, check.NotNil)
	c.Assert(err, check.IsNil)
}

func (s *S) TestValidateGrid(c *check.C) {
	err := validateGrid("03x20")
	c.Assert(err, check.IsNil)
	err = validateGrid("")
	c.Assert(err, check.NotNil)
	err = validateGrid("03a20")
	c.Assert(err, check.NotNil)
}

func (s *S) TestRotateCam(c *check.C) {
	d := dSettings{camPosition: "N"}
	d.rotateCam("D")
	c.Assert(d.camPosition, check.Equals, "L")
	d.camPosition = "S"
	d.rotateCam("E")
	c.Assert(d.camPosition, check.Equals, "L")
	d.camPosition = "N"
	d.rotateCam("E")
	c.Assert(d.camPosition, check.Equals, "O")
}

func (s *S) TestMapXYPosition(c *check.C) {
	g := grid{x: 10, y: 10}
	d := dSettings{posX: 2, posY: 2}
	mapXYPosition(d.posX, d, g)
	c.Assert(len(mappedGrid[2]), check.Equals, 2)
	d.posY = 3
	mapXYPosition(d.posX, d, g)
	fmt.Println(mappedGrid)
	c.Assert(len(mappedGrid[2]), check.Equals, 4)
}

func (s *S) TestFly(c *check.C) {
	d := dSettings{posX: 10, posY: 10, camPosition: "N"}
	g := grid{x: 20, y: 10}
	d.fly(g)
	c.Assert(d.posY, check.Equals, int64(10))
	d.camPosition = "S"
	d.fly(g)
	c.Assert(d.posY, check.Equals, int64(9))
	d.posY = 0
	d.fly(g)
	c.Assert(d.posY, check.Equals, int64(0))
	d.posX = 0
	d.camPosition = "O"
	d.fly(g)
	c.Assert(d.posX, check.Equals, int64(0))
}
