package engine

import (
	"errors"
	"math"
)

type Grid struct {
	size     Vector
	offset   Vector
	cellsize float64
	hashmap  map[int]map[int][]Placer
}

func NewGrid(size Vector, cellsize float64) *Grid {
	cellCountX := math.Ceil(size.X / cellsize)
	cellCountY := math.Ceil(size.Y / cellsize)

	hashmap := make(map[int]map[int][]Placer)

	for x := 0; x < int(cellCountX); x++ {
		hashmap[x] = make(map[int][]Placer, int(cellCountY))
		for y := 0; y < int(cellCountY); y++ {
			hashmap[x][y] = make([]Placer, 0)
		}
	}

	fullSize := Vector{
		X: cellCountX * cellsize,
		Y: cellCountY * cellsize,
	}

	offset := Vector{
		X: (size.X - fullSize.X) / 2,
		Y: (size.Y - fullSize.Y) / 2,
	}

	return &Grid{
		size:     fullSize,
		offset:   offset,
		cellsize: cellsize,
		hashmap:  hashmap,
	}
}

func (g *Grid) Add(node Placer) error {
	if g.Contains(node) == false {
		return errors.New("node outside of grid dimensions")
	}

	point := g.translatePosition(node.GetPosition())
	g.hashmap[point.X][point.Y] = append(g.hashmap[point.X][point.Y], node)

	node.Register(NodeRemoveHook, func() {
		g.remove(node)
	})

	return nil
}

func (g *Grid) GetCellOf(node Placer) ([]Placer, error) {
	if g.Contains(node) == false {
		return nil, errors.New("node outside of grid dimensions")
	}

	point := g.translatePosition(node.GetPosition())

	return g.hashmap[point.X][point.Y], nil
}

func (g *Grid) Update(node Placer) error {
	mover, ok := node.(Mover)

	if ok == false {
		return nil
	}

	currPoint := g.translatePosition(mover.GetPosition())
	lastPoint := g.translatePosition(mover.GetLastPosition())

	if currPoint.X == lastPoint.X && currPoint.Y == lastPoint.Y {
		return nil
	}

	return g.move(node, lastPoint, currPoint)
}

func (g *Grid) Contains(node Placer) bool {
	position := node.GetPosition()

	if position.X < g.offset.X || position.Y < g.offset.Y {
		return false
	}

	if position.X > (g.offset.X+g.size.X) || position.Y > (g.offset.Y+g.size.Y) {
		return false
	}

	return true
}

func (g *Grid) translatePosition(position Vector) Point {
	return Point{
		X: int(math.Floor((position.X - g.offset.X) / g.cellsize)),
		Y: int(math.Floor((position.Y - g.offset.Y) / g.cellsize)),
	}
}

func (g *Grid) move(node Placer, from, to Point) error {
	if g.Contains(node) == false {
		return errors.New("node outside of grid dimensions")
	}

	g.hashmap[from.X][from.Y] = SliceRemoveUnordered(g.hashmap[from.X][from.Y], node)
	g.hashmap[to.X][to.Y] = append(g.hashmap[to.X][to.Y], node)

	return nil
}

func (g *Grid) remove(node Placer) error {
	if g.Contains(node) == false {
		return errors.New("node outside of grid dimensions")
	}

	point := g.translatePosition(node.GetPosition())
	g.hashmap[point.X][point.Y] = SliceRemoveUnordered(g.hashmap[point.X][point.Y], node)

	return nil
}
