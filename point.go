// FlingSolver - Fling! puzzle solver
// Copyright (c) 2014, Zurab Khetsuriani (zura.khetsuriani 'at gmail.com)
// Distributed under the Boost Software License, Version 1.0.
// (See BOOST_LICENSE_1_0.txt file or copy at http://www.boost.org/LICENSE_1_0.txt)

package main

import "strconv"

type Point struct {
	X, Y int
}

func (p Point) Eq(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}
