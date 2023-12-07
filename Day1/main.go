package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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
        //result += processLine3(scanner.Text())
        result += processLine2(scanner.Text())
    }


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return result
}

func processLine(line string) (int) {

    first := ' '
    last := ' '

    for _, char := range line {

        if unicode.IsDigit(char) {
            if(first == ' ') {
                first = char
                last = char
            } else {
                last = char
            }
        }
    } 

    result := string(first) + string(last)

    res, error :=  strconv.Atoi(result)

    if error != nil {
        fmt.Println("Something went wrong")
    }

    return res
}

func processLine3(line string) (int) {

    first := ""
    last := ""
    line = strings.ReplaceAll(line, "one", "o1ne")
    line = strings.ReplaceAll(line, "two", "t2wo")
    line = strings.ReplaceAll(line, "three", "t3hree")
    line = strings.ReplaceAll(line, "four", "f4our")
    line = strings.ReplaceAll(line, "five", "f5ive")
    line = strings.ReplaceAll(line, "six", "s6ix")
    line = strings.ReplaceAll(line, "seven", "s7even")
    line = strings.ReplaceAll(line, "eight", "e8ight")
    line = strings.ReplaceAll(line, "nine", "n9ine")

    for _, char := range line {
        if unicode.IsDigit(char) {
            if(first == "") {
                first = string(char)
                last = string(char)
            } else {
                last = string(char)
            }
        }
    }

    result := first + last

    res, error :=  strconv.Atoi(result)

    if error != nil {
        fmt.Println("Something went wrong")
    }

    return res
}


func processLine2(line string) (int) {

    first := ""
    last := ""

    numbers := [9]string{"one", "two", "three","four","five", "six", "seven", "eight", "nine"}
    constring := ""

    for _, char := range line {

        constring += string(char)
        for i, s := range numbers {
            if strings.Contains(constring, s){
                if(first == "") {
                    first = strconv.Itoa(i+1)
                    constring = strings.ReplaceAll(constring, s, s[1:])
                } else{
                    last = strconv.Itoa(i+1)
                    constring = strings.ReplaceAll(constring, s, s[1:])
                }
            }
        }

        if unicode.IsDigit(char) {
            if(first == "") {
                first = string(char)
                last = string(char)
            } else {
                last = string(char)
            }
        }
    }


    result := first + last


    res, error :=  strconv.Atoi(result)

    if error != nil {
        fmt.Println("Something went wrong")
    }

    return res
}

func main()  {
    fmt.Println(readLines("./input.txt"))
}
