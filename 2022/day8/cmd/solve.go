/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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

	fmt.Println(score(scanner, part2))
}

func score(scanner *bufio.Scanner, part2 bool) int {
	var grid [][]int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		var ints []int
		for _, c := range line {
			val, _ := strconv.Atoi(c)
			ints = append(ints, val)
		}
		grid = append(grid, ints)

	}

	score := 0
	for i, row := range grid {
		for j, tree := range row {
			if part2 {
				treeScore := scenicScore(tree, i, j, grid)
				if treeScore > score {
					score = treeScore
				}
			} else {
				if treeIsVisible(tree, i, j, grid) {
					score++
				}
			}
		}
	}

	return score
}

func treeIsVisible(tree, row, col int, grid [][]int) bool {

	if row == 0 || col == 0 || row == len(grid)-1 || col == len(grid[0])-1 {
		return true
	}

	//check up
	isVisible := true
	for curr := row - 1; curr >= 0; curr-- {
		isVisible = isVisible && tree > grid[curr][col]
	}
	if isVisible {
		return true
	}
	//check down
	isVisible = true
	for curr := row + 1; curr < len(grid); curr++ {
		isVisible = isVisible && tree > grid[curr][col]
	}
	if isVisible {
		return true
	}
	//check left
	isVisible = true
	for curr := col - 1; curr >= 0; curr-- {
		isVisible = isVisible && tree > grid[row][curr]
	}
	if isVisible {
		return true
	}
	//check right
	isVisible = true
	for curr := col + 1; curr < len(grid[0]); curr++ {
		isVisible = isVisible && tree > grid[row][curr]
	}

	return isVisible
}

func scenicScore(tree, row, col int, grid [][]int) int {
	if row == 0 || col == 0 || row == len(grid)-1 || col == len(grid[0])-1 {
		return 0
	}
	score := 1
	// check up
	count := 0
	for curr := row - 1; curr >= 0; curr-- {
		count++
		if tree <= grid[curr][col] {
			break
		}
	}
	score *= count
	// check down
	count = 0
	for curr := row + 1; curr < len(grid); curr++ {
		count++
		if tree <= grid[curr][col] {
			break
		}
	}
	score *= count
	// check left
	count = 0
	for curr := col - 1; curr >= 0; curr-- {
		count++
		if tree <= grid[row][curr] {
			break
		}
	}
	score *= count
	// check right
	count = 0
	for curr := col + 1; curr < len(grid[0]); curr++ {
		count++
		if tree <= grid[row][curr] {
			break
		}
	}
	score *= count

	return score
}
