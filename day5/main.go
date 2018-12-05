package main

import (
	"fmt"
	"os"

	"github.com/urbansson/advent-of-code/util"
)

func main() {
	args := os.Args[1:]
	file := args[0]
	fc := util.ReadFile(file)[0]

	polymer := simulateReaction(fc)
	fmt.Println("polymer length", len(polymer))

}

func simulateReaction(s string) string {
	var d = 'a' - 'A'
	i := 0
	for i < len(s)-1 {
		diff := rune(s[i]) - rune(s[i+1])
		if diff == d || diff == -d {
			if i == 0 {
				s = s[2:]
			} else if i == len(s)-2 {
				s = s[:len(s)-2]
				break
			} else {
				s = s[:i] + s[i+2:]
				i--
			}
		} else {
			i++
		}
	}
	return s
}
