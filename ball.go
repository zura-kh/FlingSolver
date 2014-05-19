// FlingSolver - Fling! puzzle solver
// Copyright (c) 2014, Zurab Khetsuriani (zura.khetsuriani 'at gmail.com)
// Distributed under the Boost Software License, Version 1.0.
// (See BOOST_LICENSE_1_0.txt file or copy at http://www.boost.org/LICENSE_1_0.txt)

package main

var invalidPos = Point{-1, -1}

type Ball struct {
	pos Point
	fb  *FlingBoard
}

func (ball *Ball) Position() Point {
	return ball.pos
}

// updates board as well
func (ball *Ball) SetPosition(pos Point) {
	if pos.Eq(ball.pos) {
		return
	}

	// reset only if no other ball has been placed
	if ball.Valid() && ball.fb.Ball(ball.pos) == ball {
		ball.fb.SetBall(ball.pos, nil)
	}

	ball.pos = pos
	ball.fb.SetBall(pos, ball)
}

// the ball has valid position within the board
func (ball *Ball) Valid() bool {
	return ball.fb.ValidPosition(ball.pos)
}

// removes from the board
func (ball *Ball) invalidate() {
	ball.fb.SetBall(ball.Position(), nil)
	ball.pos = invalidPos
}

// returns the next movable ball (i.e. skips middle locked balls)
func (ball *Ball) Move(deltaPt Point, canFallOut bool) *Ball {
	nextPos := ball.Position().Add(deltaPt)

	// no move if next position is valid but not empty
	if ball.fb.ValidPosition(nextPos) && ball.fb.Ball(nextPos) != nil {
		return nil
	}

	nextPos = nextPos.Add(deltaPt)
	var nextBall *Ball = nil
	var firstCollisionFound bool = false

	for ball.fb.ValidPosition(nextPos) {
		nextBall = ball.fb.Ball(nextPos)

		if nextBall != nil {
			if !firstCollisionFound {
				ball.SetPosition(nextPos.Sub(deltaPt)) // prev pos
				firstCollisionFound = true
			}

		} else if firstCollisionFound { // and we're at empty cell
			return ball.fb.Ball(nextPos.Sub(deltaPt)) // ball at the prev pos
		}

		nextPos = nextPos.Add(deltaPt)
	}

	// we're out of the board

	// the last ball is right at the border
	if firstCollisionFound {
		return ball.fb.Ball(nextPos.Sub(deltaPt)) // ball at the prev pos (last ball)
	}

	// no collisions were found
	if canFallOut {
		ball.invalidate()
	}

	return nil
}
