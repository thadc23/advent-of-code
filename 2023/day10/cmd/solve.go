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

type Point struct {
	Y, X int
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
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	start := findStart(grid)

	findLoop(grid, start, make(map[Point]bool))

	return total
}

func findLoop(grid [][]string, current Point, visited map[Point]bool) []Point {
	directions := []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	fmt.Println(current.Y, ",", current.X, " - ", grid[current.Y][current.X])

	if visited[current] {
		return nil
	}

	visited[current] = true

	switch grid[current.Y][current.X] {
	case "S":
		if len(visited) > 1 {
			return []Point{current}
		}
		fallthrough
	case "|":
		// Vertical pipe
		next := Point{current.Y + directions[0].Y, current.X + directions[0].X}
		if next.Y >= 0 && next.Y < len(grid) && next.X >= 0 && next.X < len(grid[0]) {
			// Recursive call
			subloop := findLoop(grid, next, visited)
			if subloop != nil {
				return append([]Point{current}, subloop...)
			}
		}
	case "_":
		// horizontal pipe
		next := Point{current.Y + directions[1].Y, current.X + directions[1].X}
		if next.Y >= 0 && next.Y < len(grid) && next.X >= 0 && next.X < len(grid[0]) {
			// Recursive call
			subloop := findLoop(grid, next, visited)
			if subloop != nil {
				return append([]Point{current}, subloop...)
			}
		}
	case "7", "F", "L", "J":
		direction := directions[0]
		pipeType := grid[current.Y][current.X]

		if pipeType == "F" {
			direction = directions[1]
		} else if pipeType == "7" {
			direction = directions[2]
		} else if pipeType == "J" {
			direction = directions[3]
		}

		next := Point{current.Y + direction.Y, current.X + direction.X}

		if next.Y >= 0 && next.Y < len(grid) && next.X >= 0 && next.X < len(grid[0]) {
			// Recursive call
			subloop := findLoop(grid, next, visited)
			if subloop != nil {
				return append([]Point{current}, subloop...)
			}
		}
	}

	return nil
}

func findStart(grid [][]string) Point {
	for i, row := range grid {
		for j, col := range row {
			if col == "S" {
				return Point{i, j}
			}
		}
	}
	return Point{}
}
