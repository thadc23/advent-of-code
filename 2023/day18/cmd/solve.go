/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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
		solve(part2)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")
	solveCmd.PersistentFlags().Bool("part2", false, "Solve part2")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func solve(part2 bool) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fmt.Println(calc(scanner, part2))
}

func calc(scanner *bufio.Scanner, part2 bool) int {
	total := 0

	var grid [][]string

	grid_size := 380
	// start := 260
	for i := 0; i < grid_size; i++ {
		var row []string = make([]string, grid_size)
		for j := 0; j < grid_size; j++ {
			row[j] = "."
		}
		grid = append(grid, row)
	}

	curr_row := 254
	curr_col := 45
	grid[254][45] = "#"
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, " ")
		distance, _ := strconv.Atoi(inputs[1])

		for i := 0; i < distance; i++ {
			switch inputs[0] {
			case "U":
				// if curr_row == 0 {
				// 	var row []string = make([]string, grid_size)
				// 	row[curr_col] = "#"
				// 	grid = prependRow(grid, row)
				// } else {
				grid[curr_row-1][curr_col] = "#"
				curr_row--
				// }
			case "D":
				grid[curr_row+1][curr_col] = "#"
				curr_row++
			case "L":
				grid[curr_row][curr_col-1] = "#"
				curr_col--
			case "R":
				grid[curr_row][curr_col+1] = "#"
				curr_col++
			}
		}

	}

	for i := 0; i < grid_size; i++ {
		fmt.Println(grid[i])
	}

	for i := 0; i < grid_size; i++ {
		for j := 0; j < grid_size; j++ {
			switch grid[i][j] {
			case ".":
				if interior(grid, i, j) {
					grid[i][j] = "#"
					total++
				}
			case "#":
				total++
			}
		}
		fmt.Println(grid[i])
		// first := indexOf("#", grid[i])
		// if first == -1 {
		// 	continue
		// }
		// last := lastIndexOf("#", grid[i])
		// // fmt.Printf("First: %d Last: %d ", first, last)
		// // fmt.Println(grid[i])
		// total += last - (first - 1)
	}

	return total
}

func interior(grid [][]string, y, x int) bool {
	return left(grid[y], x) && right(grid[y], x) && up(grid, y, x) && down(grid, y, x)
}

func left(row []string, x int) bool {
	for i := x; i >= 0; i-- {
		if row[i] == "#" {
			return true
		}
	}
	return false
}
func right(row []string, x int) bool {
	for i := x; i < len(row); i++ {
		if row[i] == "#" {
			return true
		}
	}
	return false
}
func up(grid [][]string, y, x int) bool {
	for i := y; i >= 0; i-- {
		if grid[i][x] == "#" {
			return true
		}
	}
	return false
}
func down(grid [][]string, y, x int) bool {
	for i := y; i < len(grid); i++ {
		if grid[i][x] == "#" {
			return true
		}
	}
	return false
}
