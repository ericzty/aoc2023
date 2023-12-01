package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
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
		out := findNums(line)

		fmt.Println(line)
		fmt.Println(out)

		result += out[0] * 10
		result += out[len(out)-1]
	}

	fmt.Println(result)
}

func findNums(str string) (out []int) {
	for i, letter := range str {
		if num, err := strconv.Atoi(string(letter)); err == nil {
			out = append(out, num)
			continue
		}

		substr := str[i:]

		patterns := map[string]int{
			"^one":   1,
			"^two":   2,
			"^three": 3,
			"^four":  4,
			"^five":  5,
			"^six":   6,
			"^seven": 7,
			"^eight": 8,
			"^nine":  9,
			"^ten":   10,
		}

		for pattern, value := range patterns {
			match, _ := regexp.MatchString(pattern, substr)
			if match {
				out = append(out, value)
				break
			}
		}
	}
	return
}
