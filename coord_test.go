package main

import check "gopkg.in/check.v1"

func (s *S) TestStripInput(c *check.C) {
	_, err := dSettingsFromInput("")
	c.Assert(err, check.NotNil)
	d, err := dSettingsFromInput("0315ODEDEFFDFFD")
	c.Assert(len(d.coords), check.Equals, 2)
	c.Assert(d.camPosition, check.NotNil)
	c.Assert(d.commands, check.NotNil)
	c.Assert(err, check.IsNil)
}
