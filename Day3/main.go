package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func readLines(filename string) []string {
    file, err := os.ReadFile(filename)
    if err != nil {
        return nil
    }

    return strings.Split(strings.TrimSpace(string(file)), "\n")
}

func processLine2(lines []string) int {
    re := regexp.MustCompile("\\d+")

    symbols := []string{"@", "#", "$", "%", "&", "*", "/", "+", "-", "="}

    result := 0

    for i, line := range lines {
        position := re.FindAllStringIndex(line, -1)
        for _, idx := range position {

            start := idx[0]
            end := idx[1]
            startofnumber := start - 1
            endofnumber := end + 1

            if startofnumber < 0 {
                startofnumber = 0
            }
            if endofnumber > len(line) {
                endofnumber = len(line)
            }

            linestocheck := [3]string{(line)[startofnumber:endofnumber]}

            if i > 0 {
                linestocheck[1] = lines[i-1][startofnumber:endofnumber]
            }

            if i < len(lines)-1 {
                linestocheck[2] = lines[i+1][startofnumber:endofnumber]
            }

            added := false

            for _, check := range linestocheck {
                if !added {
                    for _, symbol := range symbols {
                        if strings.ContainsAny(check, symbol) {

                            number := line[start:end]
                            number2, err := strconv.Atoi(number)

                            if err != nil {
                                fmt.Println("Something went wrong")
                            }
                            result += number2
                            added = true
                            break
                        }
                    }
                }
            }
        }

    }

    return result
}

func main() {
    lines := readLines("./input.txt")
    res := 0
    res += processLine2(lines)

    fmt.Println(res)
}
