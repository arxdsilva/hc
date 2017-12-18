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

func (s *S) TestCalcMinDFromSettings(c *check.C) {
	dSett := dSettings{gridMaxX: 10, gridMaxY: 20}
	min := dSett.calcMinDFromSettings()
	c.Assert(min, check.Equals, int64(4))
	dSett = dSettings{gridMaxX: 10, gridMaxY: 9}
	min = dSett.calcMinDFromSettings()
	c.Assert(min, check.Equals, int64(3))
}
