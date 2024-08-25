package main

import (
	"flag"
	"fmt"
	"gameoflife/patterns"
	"strings"
	"time"
)

func printBoard(board [30][60]int, generation int) {
	fmt.Print("\033[H\033[2J") // Clear the screen
	for i := 0; i < 30; i++ {
		fmt.Println(join(board[i]))
	}
	fmt.Printf("generation count: %d", generation)
}

func join(line [60]int) string {
	var sb strings.Builder
	for j := 0; j < 60; j++ {
		if line[j] == 0 {
			sb.WriteString("  ")
		} else {
			sb.WriteString("* ")
		}
	}

	return sb.String()
}

func applyPattern(board [30][60]int, positions [][2]int) [30][60]int {
	for _, pos := range positions {
		board[pos[0]][pos[1]] = 1
	}
	return board
}

func neighborCount(board [30][60]int, x, y int) int {
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, direction := range directions {
		newX, newY := x+direction[0], y+direction[1]
		if newX < 0 || newX >= 30 || newY < 0 || newY >= 60 {
			continue
		}

		if board[newX][newY] == 1 {
			count++
		}
	}

	return count
}

func gameState(board [30][60]int) [30][60]int {
	newBoard := [30][60]int{}
	for i := 0; i < 30; i++ {
		for j := 0; j < 60; j++ {
			count := neighborCount(board, i, j)
			if board[i][j] == 0 {
				if count == 3 {
					newBoard[i][j] = 1
				}
			} else {
				if count < 2 || count > 3 {
					newBoard[i][j] = 0
				} else {
					newBoard[i][j] = 1
				}
			}
		}
	}

	return newBoard
}

func main() {
	board := [30][60]int{}
	speed := flag.Int("speed", 80, "render speed in millisecond")
	pattern := flag.String("pattern", "star", "choose patterns from lwss, star, glider")
	flag.Parse()

	switch *pattern {
	case "lwss":
		board = applyPattern(board, patterns.LWSS)
	case "glider":
		board = applyPattern(board, patterns.Glider)
	default:
		board = applyPattern(board, patterns.Star)
	}

	fmt.Print("\033[?25l")       // Hide cursor
	defer fmt.Print("\033[?25h") // Show cursor when the program exits
	generation := 1
	for {
		printBoard(board, generation)
		time.Sleep(time.Millisecond * time.Duration((*speed)))
		board = gameState(board)
		generation++
	}
}
