package main

import (
	"fmt"
	"strings"
	"time"
)

func printBoard(board [30][60]int) {
	fmt.Print("\033[H\033[2J") // Clear the screen
	for i := 0; i < 30; i++ {
		fmt.Println(join(board[i]))
	}
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
	// glider
	// board[0][1] = 1
	// board[1][2] = 1
	// board[2][0] = 1
	// board[2][1] = 1
	// board[2][2] = 1

	// 2 Lightweight Spaceship (2LWSS)
	board[10][11] = 1
	board[10][12] = 1
	board[10][13] = 1
	board[11][10] = 1
	board[11][14] = 1
	board[12][14] = 1
	board[13][10] = 1
	board[13][13] = 1

	board[10][21] = 1
	board[10][22] = 1
	board[10][23] = 1
	board[11][20] = 1
	board[11][24] = 1
	board[12][24] = 1
	board[13][20] = 1
	board[13][23] = 1

	// space ship
	// board[15][30] = 1
	// board[15][31] = 1
	// board[15][32] = 1
	// board[16][30] = 1
	// board[16][32] = 1
	// board[17][30] = 1
	// board[17][32] = 1

	fmt.Print("\033[?25l")       // Hide cursor
	defer fmt.Print("\033[?25h") // Show cursor when the program exits
	for {
		printBoard(board)
		time.Sleep(time.Millisecond * 80)
		board = gameState(board)
	}
}
