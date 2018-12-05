package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/urbansson/advent-of-code/util"
)

var d = 'a' - 'A'

func main() {
	args := os.Args[1:]
	file := args[0]
	fc := util.ReadFile(file)[0]

	polymer := simulateReaction(fc)
	fmt.Println("polymer length", len(polymer))

	_, length := simulateImprovedReaction(polymer)
	fmt.Println("improved polymer length", length)
}

func simulateImprovedReaction(s string) (string, int) {

	shortes := math.MaxInt64
	improvedPolymer := ""
	for i := 'a'; i <= 'z'; i++ {
		improved := strings.Replace(s, string(i), "", -1)
		improved = strings.Replace(improved, string(i-d), "", -1)
		improved = simulateReaction(improved)
		l := len(improved)
		if l < shortes {
			shortes = l
			improvedPolymer = improved
		}
	}

	return improvedPolymer, shortes
}

func simulateReaction(s string) string {
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
