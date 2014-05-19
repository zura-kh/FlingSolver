package main

import "testing"

func TestEq(t *testing.T) {
	p, q := Point{1, 2}, Point{1, 2}
	if !p.Eq(q) {
		t.Error(p, "not equals to", q)
	}

	p, q = Point{10, 5}, Point{10, -5}
	if p.Eq(q) {
		t.Error(p, "equals to", q)
	}
}

func TestAdd(t *testing.T) {
	p, q, res := Point{1, 2}, Point{-1, 10}, Point{0, 12}
	resTest := p.Add(q)
	if !resTest.Eq(res) {
		t.Error(p, "add", q, "equals to", resTest)
	}
}

func TestSub(t *testing.T) {
	p, q, res := Point{1, 2}, Point{-1, 10}, Point{2, -8}
	resTest := p.Sub(q)
	if !resTest.Eq(res) {
		t.Error(p, "sub", q, "equals to", resTest)
	}
}
