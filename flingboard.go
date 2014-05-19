// FlingSolver - Fling! puzzle solver
// Copyright (c) 2014, Zurab Khetsuriani (zura.khetsuriani 'at gmail.com)
// Distributed under the Boost Software License, Version 1.0.
// (See BOOST_LICENSE_1_0.txt file or copy at http://www.boost.org/LICENSE_1_0.txt)

package main

import "errors"

type Direction int

const (
	DirLeft Direction = iota
	DirRight
	DirUp
	DirDown
)

func (dir Direction) String() string {
	switch dir {
	case DirLeft:
		return "left"
	case DirRight:
		return "right"
	case DirUp:
		return "up"
	case DirDown:
		return "down"
	}
	return "undefined"
}

type Move struct {
	Pos Point
	Dir Direction
}

func (m Move) Eq(other Move) bool {
	return m.Pos.Eq(other.Pos) && m.Dir == other.Dir
}

type undoStateItem struct {
	ball   *Ball
	origin Point
}

type undoState []undoStateItem

type FlingBoard struct {
	width  int
	height int

	board [][]*Ball // [height][width] 2D vector for easy positioning
	balls []Ball    // keep track of all balls

	dirPointMap map[Direction]Point // dir to delta point mapping

	moves []Move
}

func NewFlingBoard(ballPositions []Point, width, height int) (*FlingBoard, error) {
	fb := new(FlingBoard)

	fb.width = width
	fb.height = height

	// set up direction <> delta point map
	fb.dirPointMap = make(map[Direction]Point)
	fb.dirPointMap[DirLeft] = Point{-1, 0}
	fb.dirPointMap[DirRight] = Point{1, 0}
	fb.dirPointMap[DirUp] = Point{0, -1}
	fb.dirPointMap[DirDown] = Point{0, 1}

	fb.balls = make([]Ball, len(ballPositions))
	for i := range fb.balls {
		fb.balls[i] = Ball{ballPositions[i], fb}
	}

	// [H][W] 2D slice (initialized to nil)
	fb.board = make([][]*Ball, height)
	for i := range fb.board {
		fb.board[i] = make([]*Ball, width)
	}

	// fill the board
	for i := range fb.balls {
		pos := fb.balls[i].Position()

		if !fb.ValidPosition(pos) {
			return nil, errors.New("Invalid ball position")
		}

		if fb.Ball(pos) != nil {
			return nil, errors.New("Duplicate ball position")
		}

		fb.SetBall(pos, &fb.balls[i])
	}
	return fb, nil
}

func (fb *FlingBoard) Width() int {
	return fb.width
}

func (fb *FlingBoard) Height() int {
	return fb.height
}

// pos is within the board
func (fb *FlingBoard) ValidPosition(pos Point) bool {
	return 0 <= pos.X && pos.X < fb.width && 0 <= pos.Y && pos.Y < fb.height
}

func (fb *FlingBoard) Ball(pos Point) *Ball {
	return fb.board[pos.Y][pos.X]
}

func (fb *FlingBoard) SetBall(pos Point, ball *Ball) {
	fb.board[pos.Y][pos.X] = ball
}

// returns moves made for solving
func (fb *FlingBoard) Moves() []Move {
	return fb.moves
}

// checks if only one ball left on the board
func (fb *FlingBoard) puzzleSolved() bool {
	ballsCount := 0
	for i := range fb.balls {
		if fb.balls[i].Valid() { // count valid balls
			ballsCount++
			if ballsCount > 1 {
				return false
			}
		}
	}

	return true
}

// main method - solves the puzzle recursively
func (fb *FlingBoard) solve() bool {
	if fb.puzzleSolved() {
		return true
	}

	for i := range fb.balls {
		if !fb.balls[i].Valid() { // skip invalid balls
			continue
		}

		// for each direction
		for dir := range fb.dirPointMap {
			uState, ok := fb.makeMove(&fb.balls[i], dir)

			if ok { // move happened
				// now try to recursively solve
				if fb.solve() {
					return true
				} else {
					fb.undoMove(uState)
				}
			}
		}
	}

	return false
}

func (fb *FlingBoard) makeMove(ball *Ball, dir Direction) (undoState, bool) {
	var deltaPt Point = fb.dirPointMap[dir]

	origin := ball.Position()

	nextBall := ball.Move(deltaPt, false)

	// the first ball should always collide
	if nextBall == nil {
		return nil, false
	}

	// append only the first ball movement
	fb.moves = append(fb.moves, Move{origin, dir})

	var uState undoState
	uState = append(uState, undoStateItem{ball, origin})

	for nextBall != nil {
		uState = append(uState, undoStateItem{nextBall, nextBall.Position()})
		nextBall = nextBall.Move(deltaPt, true)
	}

	return uState, true
}

func (fb *FlingBoard) undoMove(uState undoState) {
	for i := range uState {
		uState[i].ball.SetPosition(uState[i].origin)
	}

	fb.moves = fb.moves[0 : len(fb.moves)-1] // remove last element
}
