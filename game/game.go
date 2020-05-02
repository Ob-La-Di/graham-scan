package game

import (
	"golang.org/x/mobile/geom"
	"math/rand"
	"sort"
)

type Gift struct {
	Point      geom.Point
	IsSelected bool
}

type Game struct {
	width, height int
	Gifts         []*Gift
}

func (g *Game) Wrap() {
	pivotIndex := g.findPivotIndex()

	temp := g.Gifts[0]
	g.Gifts[0] = g.Gifts[pivotIndex]
	g.Gifts[pivotIndex] = temp

	g.SortByPivotAngle()

	stack := []*Gift{g.Gifts[0], g.Gifts[1]}

	for i := 2; i < len(g.Gifts); i++ {
		for len(stack) > 2 && cross(g.Gifts[i].Point, stack[len(stack)-1].Point, stack[len(stack)-2].Point) < 0 {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, g.Gifts[i])
	}

	for _, giftP := range stack {
		giftP.IsSelected = true
	}
}

func cross(a, b, c geom.Point) geom.Pt {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

func (g *Game) SortByPivotAngle() {
	pivot := g.Gifts[0].Point

	sort.Slice(g.Gifts[1:], func(i, j int) bool {
		a, b := g.Gifts[i+1].Point, g.Gifts[j+1].Point
		u, v := geom.Point{X: a.X - pivot.X, Y: a.Y - pivot.Y}, geom.Point{X: b.X - pivot.X, Y: b.Y - pivot.Y}

		cross := u.X*v.Y - u.Y*v.X

		return cross < 0
	})
}

func New(width, height int, numberOfPoints int) *Game {
	g := &Game{width, height, []*Gift{}}

	for i := 0; i < numberOfPoints; i++ {
		g.Gifts = append(g.Gifts, &Gift{randomPoint(width, height), false})
	}

	return g
}

func randomPt(min, max int) geom.Pt {
	return geom.Pt(rand.Intn(max-min+1) + min)
}

func randomPoint(width, height int) geom.Point {
	return geom.Point{X: randomPt(100, width-100), Y: randomPt(100, height-100)}
}

func (g Game) findPivotIndex() int {
	var pivot *Gift
	minIndex := 0

	for index, gift := range g.Gifts {
		if pivot == nil {
			pivot = gift
		} else if gift.Point.Y < pivot.Point.Y ||
			gift.Point.Y == pivot.Point.Y && gift.Point.X < pivot.Point.X {
			pivot = gift
			minIndex = index
		}
	}

	return minIndex
}
