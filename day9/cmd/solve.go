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
	"sort"
	"strconv"

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

		var input [][]int

		i := 0
		for scanner.Scan() {
			line := scanner.Text()
			input = append(input, make([]int, len(line)))
			for j, x := range line {
				val, _ := strconv.Atoi(string(x))
				input[i][j] = val
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
	result := 0
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if isLowPoint(row, col, input) {
				result += 1 + input[row][col]
			}
		}
	}

	fmt.Println(result)

}

func isLowPoint(row, col int, input [][]int) bool {

	isLow := true
	//checkUp
	if row-1 >= 0 {
		isLow = isLow && (input[row][col] < input[row-1][col])
	}
	//checkDown
	if row+1 < len(input) {
		isLow = isLow && (input[row][col] < input[row+1][col])
	}
	//checkLeft
	if col-1 >= 0 {
		isLow = isLow && (input[row][col] < input[row][col-1])
	}
	//checkRight
	if col+1 < len(input[row]) {
		isLow = isLow && (input[row][col] < input[row][col+1])
	}

	return isLow
}

func solvePart2(input [][]int) {

	result := 0
	var basins []int
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if isLowPoint(row, col, input) {
				result += 1 + input[row][col]
				basins = append(basins, findBasin(row, col, input))
			}
		}
	}

	sort.Ints(basins)
	fmt.Println(basins)

	fmt.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])

}

func findBasin(row, col int, input [][]int) int {
	count := 1
	//checkUp
	if row-1 >= 0 {
		if input[row][col] < input[row-1][col] && input[row-1][col] != 9 {
			count += findBasin(row-1, col, input)
			input[row-1][col] = 9
		}
	}
	//checkDown
	if row+1 < len(input) {
		if input[row][col] < input[row+1][col] && input[row+1][col] != 9 {
			count += findBasin(row+1, col, input)
			input[row+1][col] = 9
		}
	}
	//checkLeft
	if col-1 >= 0 {
		if input[row][col] < input[row][col-1] && input[row][col-1] != 9 {
			count += findBasin(row, col-1, input)
			input[row][col-1] = 9
		}
	}
	//checkRight
	if col+1 < len(input[row]) {
		if input[row][col] < input[row][col+1] && input[row][col+1] != 9 {
			count += findBasin(row, col+1, input)
			input[row][col+1] = 9
		}
	}
	return count
}
