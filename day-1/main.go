package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("hello world")
	fs := readFile("./input.txt")
	frequency := 0
	for _, f := range fs {
		frequency += f
	}
	fmt.Println(frequency)
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
