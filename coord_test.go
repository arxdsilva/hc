package main

import check "gopkg.in/check.v1"

func (s *S) TestStripInput(c *check.C) {
	str, err := stripInput("0315ODEDEFFDFFD")
	c.Assert(str, check.NotNil)
	c.Assert(err, check.IsNil)
}
