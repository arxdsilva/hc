package main

import (
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
