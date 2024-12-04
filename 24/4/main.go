package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Directions represent row and column changes for the 8 possible directions
var directions = [8][2]int{
	{-1, 0}, {1, 0}, // Up, Down
	{0, -1}, {0, 1}, // Left, Right
	{-1, -1}, {1, 1}, // Diagonal up-left, down-right
	{-1, 1}, {1, -1}, // Diagonal up-right, down-left
}

// Word to find
const word = "XMAS"

func main() {
	// Read the word search from input.txt
	grid, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Solve Part 1: Find occurrences of "XMAS"
	part1Count := findOccurrences(grid, word)
	fmt.Printf("Part 1: The word '%s' appears %d times in the word search.\n", word, part1Count)

	// Solve Part 2: Find occurrences of X-MAS patterns
	part2Count := findXMASCounts(grid)
	fmt.Printf("Part 2: The X-MAS pattern appears %d times in the word search.\n", part2Count)
}

// readInput reads the grid from a file
func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

// findOccurrences counts occurrences of the word in all directions
func findOccurrences(grid [][]rune, word string) int {
	wordLen := len(word)
	wordRunes := []rune(word)
	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			// Check all 8 directions
			for _, dir := range directions {
				if matchWord(grid, row, col, dir, wordRunes, wordLen) {
					count++
				}
			}
		}
	}
	return count
}

// matchWord checks if the word matches in the given direction
func matchWord(grid [][]rune, row, col int, dir [2]int, word []rune, wordLen int) bool {
	for i := 0; i < wordLen; i++ {
		r, c := row+i*dir[0], col+i*dir[1]
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != word[i] {
			return false
		}
	}
	return true
}

// findXMASCounts checks for the X-MAS pattern as described
func findXMASCounts(grid [][]rune) int {
	count := 0

	// Iterate through the grid, excluding the edges
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			// If the character is 'A', check its diagonals
			if grid[row][col] == 'A' && isXMAS(grid, row, col) {
				count++
			}
		}
	}
	return count
}

// isXMAS checks if an 'A' at (row, col) forms an X-MAS pattern
func isXMAS(grid [][]rune, row, col int) bool {
	// Get diagonal characters
	topLeft := grid[row-1][col-1]
	topRight := grid[row-1][col+1]
	bottomLeft := grid[row+1][col-1]
	bottomRight := grid[row+1][col+1]

	// Check diagonals for both 'M' and 'S' in any order
	diagonal1 := []rune{topLeft, bottomRight}
	diagonal2 := []rune{topRight, bottomLeft}

	return containsMAS(diagonal1) && containsMAS(diagonal2)
}

// containsMAS checks if the slice contains both 'M' and 'S' in any order
func containsMAS(diagonal []rune) bool {
	mFound, sFound := false, false
	for _, char := range diagonal {
		if char == 'M' {
			mFound = true
		} else if char == 'S' {
			sFound = true
		}
	}
	return mFound && sFound
}
