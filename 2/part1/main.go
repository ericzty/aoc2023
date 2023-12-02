package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	result := 0

	str := ""
	if file, err := ioutil.ReadFile("input"); err != nil {
		panic(err)
	} else {
		str = string(file)
	}

	scanner := bufio.NewScanner(strings.NewReader(str))

	for scanner.Scan() {
		line := scanner.Text()

		id, isValid := counter(line)

		if isValid {
			result += id
		}

	}

	fmt.Println(result)
}

func counter(str string) (id int, isValid bool) {
	fmt.Println(str)

	splitGame := strings.Split(str, ":")
	id, _ = strconv.Atoi(string(splitGame[0][5:]))
	fmt.Println("id: ", id)

	splitRounds := strings.Split(splitGame[1], ";")
	red, green, blue := 12, 13, 14
	gameRed, gameGreen, gameBlue := 0, 0, 0
	for _, round := range splitRounds {
		fmt.Println("round: ", round)
		splitColors := strings.Split(round, ",")
		for _, color := range splitColors {
			if strings.Contains(color, "red") {
				red, _ := strconv.Atoi(strings.Split(color, " ")[1])
				if red > gameRed {
					gameRed = red
				}
			}
			if strings.Contains(color, "green") {
				green, _ := strconv.Atoi(strings.Split(color, " ")[1])
				if green > gameGreen {
					gameGreen = green
				}
			}
			if strings.Contains(color, "blue") {
				blue, _ := strconv.Atoi(strings.Split(color, " ")[1])
				if blue > gameBlue {
					gameBlue = blue
				}
			}
		}
	}
	fmt.Println("max rgb: ", gameRed, gameGreen, gameBlue)
	if gameRed > red || gameGreen > green || gameBlue > blue {
		fmt.Println("Invalid")
		isValid = false
		return
	}

	fmt.Println("Valid")
	isValid = true
	return
}
