package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fs := readFile("./input.txt")
	pm := make(map[int]bool)
	frequency := 0
	done := false
	for !done {
		for _, f := range fs {
			frequency += f
			_, ok := pm[frequency]
			if ok {
				fmt.Println("Already have:", frequency)
				done = true
				break
			} else {
				pm[frequency] = true
			}
		}
	}
}

func readFile(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		s = append(s, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return s
}
