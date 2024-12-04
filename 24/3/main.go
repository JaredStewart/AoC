package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// Regular expression to match valid mul(X,Y) instructions
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	totalSum := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// Extract numbers from the regex groups
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])

			// Multiply and add to the total sum
			totalSum += x * y
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Output the total sum
	fmt.Println("Total sum of valid multiplications:", totalSum)
}

func main() {
	// Read the entire input into a buffer
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	input := string(data)

	totalSum := 0
	mulEnabled := true // Initially, multiplications are enabled
	buffer := strings.Builder{}
	state := "none" // Tracks the current state: "none", "mul", "do", "dont"

	// Walk through the input character by character
	for _, char := range input {
		switch state {
		case "none":
			// Look for the start of an instruction
			if char == 'm' {
				state = "maybe_mul"
				buffer.WriteRune(char)
			} else if char == 'd' {
				state = "maybe_do_or_dont"
				buffer.WriteRune(char)
			}
		case "maybe_mul":
			// Build the potential mul instruction
			buffer.WriteRune(char)
			if buffer.String() == "mul(" {
				state = "in_mul"
				buffer.Reset()
			} else if !strings.HasPrefix("mul(", buffer.String()) {
				state = "none"
				buffer.Reset()
			}
		case "in_mul":
			// Process the numbers in the mul instruction
			if char == ')' {
				parts := strings.Split(buffer.String(), ",")
				if len(parts) == 2 {
					x, err1 := strconv.Atoi(parts[0])
					y, err2 := strconv.Atoi(parts[1])
					if err1 == nil && err2 == nil && mulEnabled {
						totalSum += x * y
					}
				}
				state = "none"
				buffer.Reset()
			} else if char >= '0' && char <= '9' || char == ',' {
				buffer.WriteRune(char)
			} else {
				// Invalid character encountered; reset state
				state = "none"
				buffer.Reset()
			}
		case "maybe_do_or_dont":
			// Build the potential do() or don't() instruction
			buffer.WriteRune(char)
			if buffer.String() == "do()" {
				mulEnabled = true
				state = "none"
				buffer.Reset()
			} else if buffer.String() == "don't()" {
				mulEnabled = false
				state = "none"
				buffer.Reset()
			} else if !strings.HasPrefix("do(", buffer.String()) && !strings.HasPrefix("don't(", buffer.String()) {
				state = "none"
				buffer.Reset()
			}
		}
	}

	// Output the total sum
	fmt.Println("Total sum of valid and enabled multiplications:", totalSum)
}
