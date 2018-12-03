package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urbansson/advent-of-code/util"
)

type claim struct {
	id   int
	x, y int
	w, h int

	overlaps bool
}

func main() {
	args := os.Args[1:]
	file := args[0]
	fmt.Println("Using input:", file)
	fc := util.ReadFile(file)

	// Create the grid
	grid := make([][][]*claim, 1000)
	for y := range grid {
		grid[y] = make([][]*claim, 1000)
		for x := range grid[y] {
			grid[y][x] = make([]*claim, 0)
		}
	}

	// Parse the input
	claims := make([]*claim, len(fc))
	for i, f := range fc {
		claims[i] = parse(f)
	}

	// "Render" th
	for _, claim := range claims {
		maxX := claim.x + claim.w
		maxY := claim.y + claim.h
		for y := claim.y; y < maxY; y++ {
			for x := claim.x; x < maxX; x++ {
				grid[y][x] = append(grid[y][x], claim)
			}
		}
	}

	total := 0
	for y, row := range grid {
		for x, claims := range row {
			claimCount := len(claims)
			if claimCount > 1 {
				total++
				for _, c := range grid[y][x] {
					c.overlaps = true
				}
			}
		}
	}
	fmt.Println(total)
	for _, c := range claims {
		if !c.overlaps {
			fmt.Println(c.id)
		}
	}

	visualise(grid)
}

var replacer = strings.NewReplacer("#", "", "@ ", "", ":", "")

func parse(l string) *claim {
	cleaned := replacer.Replace(l)
	s := strings.Split(cleaned, " ")
	cords := strings.Split(s[1], ",")
	size := strings.Split(s[2], "x")

	id, _ := strconv.Atoi(s[0])
	x, _ := strconv.Atoi(cords[0])
	y, _ := strconv.Atoi(cords[1])
	w, _ := strconv.Atoi(size[0])
	h, _ := strconv.Atoi(size[1])

	return &claim{id, x, y, w, h, false}
}

func visualise(grid [][][]*claim) {
	f, _ := os.Create("/tmp/out.txt")
	defer f.Close()

	for _, row := range grid {
		r := ""
		for _, claims := range row {
			count := len(claims)
			if count == 0 {
				r += "."
			} else {
				r += strconv.Itoa(count)
			}
		}
		f.WriteString(r + "\n")
	}
}
