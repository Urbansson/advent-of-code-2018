package main

import (
	"fmt"
	"os"

	"github.com/urbansson/advent-of-code/util"
)

func main() {
	args := os.Args[1:]
	file := args[0]
	fmt.Println("Using input:", file)
	fc := util.ReadFile(file)
	twoTimes := 0
	treeTimes := 0
	for _, f := range fc {
		if hasOccurrences(f, 2) {
			twoTimes++
		}
		if hasOccurrences(f, 3) {
			treeTimes++
		}
	}
	fmt.Println(twoTimes * treeTimes)
}

func hasOccurrences(row string, occurrences int) bool {
	for _, char1 := range row {
		count := 0
		for _, char2 := range row {
			if char1 == char2 {
				count++
			}
		}
		if count == occurrences {
			return true
		}
	}
	return false
}
