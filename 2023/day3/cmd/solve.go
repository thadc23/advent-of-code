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
	var grid [][]string
	for scanner.Scan() {
		line := scanner.Text()
		row := parseRow(line)
		grid = append(grid, row)
	}
	if !part2 {
		return processGrid(grid)
	}
	return processGridPart2(grid)
}

func parseRow(line string) []string {
	var row []string
	for _, char := range line {
		row = append(row, string(char))
	}
	return row
}

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

func processGrid(grid [][]string) int {
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "." {
				continue
			}
			if !isDigit(grid[i][j]) {
				// If the current character is not a digit, check all adjacent spaces
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[i]) && isDigit(grid[x][y]) {
							// If adjacent space contains a digit, process the number
							total += processNumber(grid, x, y)
						}
					}
				}
			}
		}
	}
	return total
}
func processGridPart2(grid [][]string) int {
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "." {
				continue
			}
			if grid[i][j] == "*" {
				var nums []int
				// If the current character is not a digit, check all adjacent spaces
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[i]) && isDigit(grid[x][y]) {
							// If adjacent space contains a digit, process the number
							nums = append(nums, processNumber(grid, x, y))
						}
					}
				}
				if len(nums) == 2 {
					total += nums[0] * nums[1]
				}
			}
		}
	}
	return total
}

func processNumber(grid [][]string, row, col int) int {
	total := toInt(grid[row][col])
	grid[row][col] = "."
	multiplier := 10

	//check right
	for y := col + 1; y < len(grid[row]); y++ {
		if !isDigit(grid[row][y]) {
			break
		}
		total = total*10 + toInt(grid[row][y])
		multiplier *= 10
		grid[row][y] = "."
	}

	//check Left
	for y := col - 1; y >= 0; y-- {
		if !isDigit(grid[row][y]) {
			break
		}
		total = toInt(grid[row][y])*multiplier + total
		multiplier *= 10
		grid[row][y] = "."
	}

	return total
}

func toInt(x string) int {
	i, _ := strconv.Atoi(x)
	return i
}
