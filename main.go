// FlingSolver - Fling! puzzle solver
// Copyright (c) 2014, Zurab Khetsuriani (zura.khetsuriani 'at gmail.com)
// Distributed under the Boost Software License, Version 1.0.
// (See BOOST_LICENSE_1_0.txt file or copy at http://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"fmt"
)

func main() {

	{
		var ballPositions []Point
		ballPositions = append(ballPositions, Point{0, 0})

		fb, err := NewFlingBoard(ballPositions, 1, 1)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			b := fb.solve()
			fmt.Println(b)
			fmt.Println(fb.Moves())
		}
	}

	fmt.Println("------------")

	{
		var ballPositions []Point
		ballPositions = append(ballPositions, Point{4, 0})
		ballPositions = append(ballPositions, Point{0, 1})
		ballPositions = append(ballPositions, Point{5, 2})
		ballPositions = append(ballPositions, Point{1, 3})
		ballPositions = append(ballPositions, Point{3, 3})
		ballPositions = append(ballPositions, Point{4, 3})
		ballPositions = append(ballPositions, Point{3, 4})
		ballPositions = append(ballPositions, Point{4, 4})
		ballPositions = append(ballPositions, Point{2, 5})
		ballPositions = append(ballPositions, Point{3, 5})
		ballPositions = append(ballPositions, Point{2, 6})
		ballPositions = append(ballPositions, Point{4, 6})
		ballPositions = append(ballPositions, Point{2, 7})

		fb, err := NewFlingBoard(ballPositions, 7, 8)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			b := fb.solve()
			fmt.Println(b)
			fmt.Println(fb.Moves())
		}
	}
}
