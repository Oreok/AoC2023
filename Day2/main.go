package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
		result += processLine2(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func processLine2(line string) int {

	re := regexp.MustCompile("[0-9]+")
	line = strings.ReplaceAll(line, ";", ",")
	line = strings.ReplaceAll(line, ":", ",")

	stringArray := strings.Split(line, ",")

    red := 0
    blue := 0
    green := 0

	for _, item := range stringArray {

		number := re.FindAllString(item, -1)

		res, err := strconv.Atoi(number[0])

		if err != nil {
			fmt.Println("Something went wrong")
		}

		if strings.Contains(item, "red") && res > red {
			red = res
		} else if strings.Contains(item, "green") && res > green {
            green = res
		} else if strings.Contains(item, "blue") && res > blue {
            blue = res
		}
	}
    result := red*blue*green
    return result
}
func processLine(line string) int {

	re := regexp.MustCompile("[0-9]+")
	line = strings.ReplaceAll(line, ";", ",")
	line = strings.ReplaceAll(line, ":", ",")

	stringArray := strings.Split(line, ",")

	gameId := re.FindAllString(stringArray[0], -1)

	for _, item := range stringArray {

		number := re.FindAllString(item, -1)

		res, err := strconv.Atoi(number[0])

		if err != nil {
			fmt.Println("Something went wrong")
		}

		if strings.Contains(item, "red") && res > 12 {
			return 0
		} else if strings.Contains(item, "green") && res > 13 {
			return 0
		} else if strings.Contains(item, "blue") && res > 14 {
			return 0
		}

	}

	res, err := strconv.Atoi(gameId[0])

	if err != nil {
		fmt.Println("Something went wrong")
	}

	return res

}

func main()  {
    fmt.Println(readLines("./input.txt"))
}
