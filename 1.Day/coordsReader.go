package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Data
	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	// File handling
	filepath := "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var calibration int = 0

	// Loop for reading coords lines
	for idk, line := range lines {
		reg_two_digits := regexp.MustCompile("^.*?(one|two|three|four|five|six|seven|eight|nine|\\d).*(one|two|three|four|five|six|seven|eight|nine|\\d)")
		matches := reg_two_digits.FindStringSubmatch(line)
		if len(matches) == 0 {
			reg_one_digit := regexp.MustCompile(".*(\\d).*")
			matches = reg_one_digit.FindStringSubmatch(line)
		}

		if len(matches) > 0 {
			first := numberMap[matches[1]]
			last := numberMap[matches[len(matches)-1]]
			var number_from_line string = first + last
			fmt.Printf("%d. %s\n", idk+1, number_from_line)
			num, err := strconv.Atoi(number_from_line)
			if err != nil {
				fmt.Println("Atoi reading error", err)
				return
			}
			calibration += num
		}
	}

	fmt.Println("Calibration value is ", calibration)
}
