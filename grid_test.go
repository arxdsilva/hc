package main

import check "gopkg.in/check.v1"

func (s *S) TestSetGrid(c *check.C) {
	g := grid{}
	err := g.setGrid("")
	c.Assert(err, check.NotNil)
	err = g.setGrid("03x10")
	c.Assert(err, check.IsNil)
	c.Assert(g.x, check.NotNil)
	c.Assert(g.y, check.NotNil)
}
