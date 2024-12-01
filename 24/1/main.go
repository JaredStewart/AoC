package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var inputFile string

	// Root command setup
	rootCmd := &cobra.Command{
		Use:   "processfile",
		Short: "Processes a file with pairs of integers",
		RunE: func(cmd *cobra.Command, args []string) error {
			return processFile(inputFile)
		},
	}

	// Adding a flag for the input file
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "input.txt", "Input file to read integers from")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func processFile(inputFile string) error {
	partOne(inputFile)
	partTwo(inputFile)
	return nil
}

func partOne(inputFile string) error {
	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", inputFile, err)
	}
	defer file.Close()

	var leftInts, rightInts []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split the line into integers
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return fmt.Errorf("invalid line format: %s", line)
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("failed to parse integer: %s", parts[0])
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("failed to parse integer: %s", parts[1])
		}

		leftInts = append(leftInts, left)
		rightInts = append(rightInts, right)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Sort the slices independently
	sort.Ints(leftInts)
	sort.Ints(rightInts)

	// Calculate the running total of absolute differences
	if len(leftInts) != len(rightInts) {
		return fmt.Errorf("mismatch in the number of integers between left and right slices")
	}

	var runningTotal int
	for i := 0; i < len(leftInts); i++ {
		diff := math.Abs(float64(leftInts[i] - rightInts[i]))
		runningTotal += int(diff)
	}

	// Output the result
	fmt.Println("Running Total:", runningTotal)

	return nil
}

func partTwo(inputFile string) error {
	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", inputFile, err)
	}
	defer file.Close()

	var leftInts []int
	rightInts := make(map[int]int)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split the line into integers
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return fmt.Errorf("invalid line format: %s", line)
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("failed to parse integer: %s", parts[0])
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("failed to parse integer: %s", parts[1])
		}

		leftInts = append(leftInts, left)
		if _, ok := rightInts[right]; !ok {
			rightInts[right] = 0
		}
		if _, ok := rightInts[left]; !ok {
			rightInts[left] = 0
		}
		rightInts[right] += 1
	}

	score := 0
	for _, i := range leftInts {
		score += i * rightInts[i]
	}

	fmt.Println("Similarity Score:", score)

	return nil
}
