package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/urbansson/advent-of-code/util"
)

type point struct {
	x, y     int
	area     int
	infinite bool
}

type grid struct {
	minX, minY int
	maxX, maxY int
}

func (g *grid) updateSize(p point) {
	if p.x > g.maxX {
		g.maxX = p.x
	}
	if p.y > g.maxY {
		g.maxY = p.y
	}

	if p.x < g.minX {
		g.minX = p.x
	}
	if p.y < g.minY {
		g.minY = p.y
	}
}

func (g grid) calculateDangerAreas(points []*point) {
	for x := g.minX; x <= g.maxX; x++ {
		for y := g.minY; y <= g.maxY; y++ {
			var index int
			dist := math.MaxInt64
			dup := false
			for i, p := range points {
				d := distance(*p, point{x, y, 0, false})
				if d == 0 {
					dist = d
					index = i
					break
				} else if d < dist {
					dist = d
					dup = false
					index = i
				} else if d == dist {
					dup = true
				}
			}

			if !dup {
				p := points[index]
				// If pint is in one of the corners of the grid it will be inf
				if p.x == g.minX || p.x == g.maxX || p.y == g.minY || p.y == g.maxY {
					p.infinite = true
				}
				p.area++
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	file := args[0]
	fmt.Println("Using input:", file)
	fc := util.ReadFile(file)

	var points []*point
	var grid = &grid{
		minX: math.MaxInt64,
		minY: math.MaxInt64,
	}

	for _, f := range fc {
		s := strings.Split(f, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(strings.TrimSpace(s[1]))

		point := point{
			x: int(x),
			y: int(y),
		}
		points = append(points, &point)

		grid.updateSize(point)
	}

	grid.calculateDangerAreas(points)

	var max = &point{}
	for _, p := range points {
		if !p.infinite && max.area < p.area {
			max = p
		}
	}
	fmt.Println("Larges area is:", max.area)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(a, b point) int {
	return abs((a.x)-b.x) + abs((a.y)-b.y)
}
