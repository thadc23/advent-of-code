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

	"github.com/spf13/cobra"
)

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

		input := make([][]int, 10)

		i := 0
		for scanner.Scan() {
			input[i] = make([]int, 10)
			line := scanner.Text()
			for j, x := range line {
				input[i][j] = int(x - '0')
			}
			i++
		}

		if part2 {
			solvePart2(input)
		} else {
			solvePart1(input)
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

func solvePart1(input [][]int) {
	for i := 0; i < 100; i++ {
		input = increaseCount(input)
		for i, _ := range input {
			for j, val := range input[i] {
				if val == 10 {
					flash(i, j, input)
				}
			}
		}
		input, _ = reset(input)
	}

	fmt.Println(flashes)
}
func solvePart2(input [][]int) {
	i := 0
	done := false
	for true {
		input = increaseCount(input)
		for i, _ := range input {
			for j, val := range input[i] {
				if val == 10 {
					flash(i, j, input)
				}
			}
		}
		input, done = reset(input)
		if done {
			fmt.Println(i + 1)
			break
		}
		i++
	}

}

func increaseCount(input [][]int) [][]int {
	for i, _ := range input {
		for j, _ := range input[i] {
			input[i][j]++
		}
	}
	return input
}

var flashes int

func flash(row, col int, input [][]int) [][]int {

	flashes++
	input[row][col] = -1

	//Up
	if row-1 >= 0 {
		if input[row-1][col] >= 0 {
			input[row-1][col]++
		}
		if input[row-1][col] > 9 {
			input = flash(row-1, col, input)
		}

		//UpLeft
		if col-1 >= 0 {
			if input[row-1][col-1] >= 0 {
				input[row-1][col-1]++
			}
			if input[row-1][col-1] > 9 {
				input = flash(row-1, col-1, input)
			}
		}

		//UpRight
		if col+1 < len(input[row]) {
			if input[row-1][col+1] >= 0 {
				input[row-1][col+1]++
			}
			if input[row-1][col+1] > 9 {
				input = flash(row-1, col+1, input)
			}
		}
	}
	//checkDown
	if row+1 < len(input) {
		if input[row+1][col] >= 0 {
			input[row+1][col]++
		}
		if input[row+1][col] > 9 {
			input = flash(row+1, col, input)
		}

		//DownLeft
		if col-1 >= 0 {
			if input[row+1][col-1] >= 0 {
				input[row+1][col-1]++
			}
			if input[row+1][col-1] > 9 {
				input = flash(row+1, col-1, input)
			}
		}

		//DownRight
		if col+1 < len(input[row]) {
			if input[row+1][col+1] >= 0 {
				input[row+1][col+1]++
			}
			if input[row+1][col+1] > 9 {
				input = flash(row+1, col+1, input)
			}
		}
	}

	//checkLeft
	if col-1 >= 0 {
		if input[row][col-1] >= 0 {
			input[row][col-1]++
		}
		if input[row][col-1] > 9 {
			input = flash(row, col-1, input)
		}
	}
	//checkRight
	if col+1 < len(input[row]) {
		if input[row][col+1] >= 0 {
			input[row][col+1]++
		}
		if input[row][col+1] > 9 {
			input = flash(row, col+1, input)
		}
	}
	return input
}

func reset(input [][]int) ([][]int, bool) {
	count := 0
	for i, _ := range input {
		for j, val := range input[i] {
			if val == -1 {
				input[i][j] = 0
				count++
			}
		}
	}
	return input, count == 100
}
