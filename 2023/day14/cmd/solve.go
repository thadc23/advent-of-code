/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

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

	var hist []string

	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	if !part2 {
		north(grid)
	} else {
		target := 1000000000
		for i := 1; i <= target; i++ {
			now := cycle(grid)

			first_occurence := indexOf(now, hist)
			if first_occurence != -1 {
				cycle_length := i - first_occurence
				remaining_cycles := target - i
				final_state := first_occurence + (remaining_cycles % cycle_length)
				grid = makeGrid(hist[final_state-1], len(grid[0]))
				break
			}
			hist = append(hist, now)
		}
	}

	for i, row := range grid {
		count := 0
		for _, col := range row {
			if col == "O" {
				count++
			}
		}
		total += (len(grid) - i) * count
	}

	return total
}

func makeGrid(input string, chunkSize int) [][]string {

	var result [][]string

	for len(input) > 0 {
		// Determine the actual chunk size, considering UTF-8 encoding
		size := utf8.RuneCountInString(input)
		if size > chunkSize {
			size = chunkSize
		}

		// Extract the chunk and append it to the result
		chunk := input[:size]
		result = append(result, strings.Split(chunk, ""))

		// Remove the processed chunk from the input
		input = input[size:]
	}
	return result

}

func indexOf(val string, slice []string) int {
	for i, v := range slice {
		if v == val {
			return i + 1
		}
	}
	return -1
}

func cycle(grid [][]string) string {
	north(grid)
	west(grid)
	south(grid)
	east(grid)

	retval := ""

	for _, row := range grid {
		retval += strings.Join(row, "")
	}
	return retval
}

func north(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "O" {
				y := i - 1

				for y >= 0 && grid[y][j] == "." {
					grid[y][j] = "O"
					grid[y+1][j] = "."
					y--
				}

			}
		}
	}
}
func south(grid [][]string) {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "O" {
				y := i + 1

				for y < len(grid) && grid[y][j] == "." {
					grid[y][j] = "O"
					grid[y-1][j] = "."
					y++
				}

			}
		}
	}
}
func west(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "O" {
				x := j - 1

				for x >= 0 && grid[i][x] == "." {
					grid[i][x] = "O"
					grid[i][x+1] = "."
					x--
				}

			}
		}
	}
}
func east(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] == "O" {
				x := j + 1

				for x < len(grid[i]) && grid[i][x] == "." {
					grid[i][x] = "O"
					grid[i][x-1] = "."
					x++
				}

			}
		}
	}
}
