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

	for p, l1 := range fc {
		for _, l2 := range fc[p+1:] {
			// If same skip
			if l1 == l2 {
				continue
			}
			count := 0
			index := 0
			for cp, char1 := range l1 {
				//char not same at pos cp
				if []rune(l2)[cp] != char1 {
					count++
					index = cp
				}
			}
			if count == 1 {
				fmt.Println(l1[:index] + l1[index+1:])
				return
			}
		}
	}

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
