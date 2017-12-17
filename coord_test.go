package main

import (
	check "gopkg.in/check.v1"
)

func (s *S) TestStripInput(c *check.C) {
	dSett := dSettings{}
	err := dSett.dSettingsFromInput("")
	c.Assert(err, check.NotNil)
	err = dSett.dSettingsFromInput("0315ODEDEFFDFFD")
	c.Assert(len(dSett.coords), check.Equals, 2)
	c.Assert(dSett.camPosition, check.NotNil)
	c.Assert(dSett.commands, check.NotNil)
	c.Assert(err, check.IsNil)
}

func (s *S) TestSetGrid(c *check.C) {
	dSett := dSettings{}
	err := dSett.setGrid("")
	c.Assert(err, check.NotNil)
	err = dSett.setGrid("03x10")
	c.Assert(err, check.IsNil)
	c.Assert(len(dSett.grid), check.Equals, 2)
}

func (s *S) TestValidateGrid(c *check.C) {
	valid, err := validateGrid("03x20")
	c.Assert(err, check.IsNil)
	c.Assert(valid, check.Equals, true)
	valid, err = validateGrid("")
	c.Assert(err, check.NotNil)
	c.Assert(valid, check.Equals, false)
	valid, err = validateGrid("03a20")
	c.Assert(err, check.IsNil)
	c.Assert(valid, check.Equals, false)
}
