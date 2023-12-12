package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

func readLines(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		result += processLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func processLine(line string) int {

	re := regexp.MustCompile("[0-9]+")
	res := 0

	lines := strings.Split(line, "|")
	_, myNum, _ := strings.Cut(lines[0], ":")

	myNumArray := re.FindAllString(myNum, -1)
	myWinArray := re.FindAllString(lines[1], -1)

	for _, num := range myNumArray {

		if slices.Contains(myWinArray, num) {
			if res == 0 {
				res = 1
			} else {
                res *= 2
            }
		}

	}
	return res
}

func main() {
	fmt.Println(readLines("./input.txt"))
}
