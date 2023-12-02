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

		a := counter(line)
		result += a

	}

	fmt.Println(result)
}

func counter(str string) (result int) {
	fmt.Println(str)

	splitGame := strings.Split(str, ":")

	splitRounds := strings.Split(splitGame[1], ";")
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
	return gameRed * gameGreen * gameBlue

	return
}
