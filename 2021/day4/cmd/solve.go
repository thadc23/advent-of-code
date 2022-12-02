/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Board [][]string

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		part2, _ := cmd.Flags().GetBool("part2")

		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		if part2 {
			solvePart2(scanner)
		} else {
			solvePart1(scanner)
		}
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	solveCmd.PersistentFlags().Bool("part2", false, "Solve part2")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var boards []Board

func solvePart1(scanner *bufio.Scanner) {
	numbers := buildBoards(scanner)

	for _, num := range numbers {
		winner, board, _ := call(num)
		if winner {
			fmt.Println(calculateScore(num, board))
			return
		}
	}
}

func solvePart2(scanner *bufio.Scanner) {
	numbers := buildBoards(scanner)
	for _, num := range numbers {
		done, board := callAllWinners(num)
		if done {
			fmt.Println(calculateScore(num, board))
			return
		}
	}
}

func buildBoards(scanner *bufio.Scanner) []string {

	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")
	scanner.Scan()

	row := 0
	board := make(Board, 5)
	for scanner.Scan() {
		line := scanner.Text()
		if row == 5 {
			row = 0
			boards = append(boards, board)
			board = make(Board, 5)
		} else {
			board[row] = strings.Fields(line)
			row++
		}
	}
	return numbers
}

func call(num string) (bool, Board, int) {
	for idx, board := range boards {
		row := 0
		for row < 5 {
			for col, val := range board[row] {
				if val == num {
					board[row][col] = "X"
					if isAWinner(board, row, col) {
						return true, board, idx
					}
				}
			}
			row++
		}
	}
	return false, nil, -1
}

func callAllWinners(num string) (bool, Board) {
	done := false
	for !done {
		winner, board, idx := call(num)
		if winner && len(boards) != 1 {
			boards = remove(boards, idx)
		} else if winner && len(boards) == 1 {
			return true, board
		}
		done = !winner
	}
	return false, nil
}

func isAWinner(board Board, row, col int) bool {
	return checkRow(row, board) || checkCol(col, board)
}

func checkRow(row int, board Board) bool {
	col := 0
	for col < 5 {
		if board[row][col] != "X" {
			return false
		}
		col++
	}
	return true
}

func checkCol(col int, board Board) bool {
	row := 0
	for row < 5 {
		if board[row][col] != "X" {
			return false
		}
		row++
	}
	return true
}

func calculateScore(num string, board Board) int {
	fmt.Println(board)
	fmt.Println(num)
	score := 0
	row := 0
	for row < 5 {
		col := 0
		for col < 5 {
			if board[row][col] != "X" {
				value, _ := strconv.Atoi(board[row][col])
				score += value
			}
			col++
		}
		row++
	}
	numVal, _ := strconv.Atoi(num)
	return score * numVal
}

func remove(slice []Board, s int) []Board {
	return append(slice[:s], slice[s+1:]...)
}
