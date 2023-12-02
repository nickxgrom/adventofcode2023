package days

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Trebuchet - Day 1: https://adventofcode.com/2023/day/1
func Trebuchet() {
	f, err := os.Open("assets/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		//Part 1
		sum1 += getCalibrationValue(scanner.Text())
		//Part 2
		sum2 += wordsToDigit(scanner.Text())
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func getCalibrationValue(str string) int {
	d1 := -1
	d2 := -1

	for i, c := range str {
		if d1 < 0 && c >= 48 && c <= 57 {
			d1 = int(c - 48)
		}

		lc := str[len(str)-i-1]

		if d2 < 0 && lc >= 48 && lc <= 57 {
			d2 = int(lc - 48)
		}

		if d1 >= 0 && d2 >= 0 {
			break
		}
	}

	return d1*10 + d2
}

func wordsToDigit(str string) int {
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	hashMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	d1 := -1
	d2 := -1
	for index := range str {
		s1 := str[0:index]
		s2 := str[len(str)-index:]

		for _, d := range digits {
			i1 := strings.Index(s1, d)
			i2 := strings.Index(s2, d)

			if d1 < 0 && i1 > -1 {
				d1 = hashMap[d]
			}

			if d2 < 0 && i2 > -1 {
				d2 = hashMap[d]
			}

			if d1 > -1 && d2 > -1 {
				break
			}

		}
		if d1 != -1 && d2 != -1 {
			break
		}
	}

	if d1 < 0 || d2 < 0 {
		d := math.Max(float64(d1), float64(d2))
		return int(d*10 + d)
	}

	return d1*10 + d2
}
