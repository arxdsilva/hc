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

func (s *S) TestRegisterD(c *check.C) {
	g := grid{}
	g.setGrid("10x10")
	d := dSettings{posX: 12, posY: 1}
	err := g.registerD(d)
	c.Assert(err, check.NotNil)
	d.posX = 1
	err = g.registerD(d)
	c.Assert(err, check.IsNil)
	err = g.registerD(d)
	c.Assert(err, check.NotNil)
	d.posY = 2
	err = g.registerD(d)
	c.Assert(err, check.IsNil)
	c.Assert(len(g.droneStartPoint[1]), check.Equals, 2)
}
