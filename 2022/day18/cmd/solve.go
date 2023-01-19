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

	var grid [20][20][20]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		z, _ := strconv.Atoi(line[2])

		grid[x][y][z] = "X"
	}

	total := 0
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if grid[x][y][z] == "X" {
					total += calculateSides(x, y, z, grid)
				} else if part2 && isInterior(x, y, z, grid) {
					total -= (6 - calculateSides(x, y, z, grid))
				}
			}
		}
	}

	return total
}

func calculateSides(x, y, z int, grid [20][20][20]string) int {

	total := 6
	//check X
	if x != 0 && grid[x-1][y][z] == "X" {
		total--
	}
	if x != len(grid)-1 && grid[x+1][y][z] == "X" {
		total--
	}
	//check Y
	if y != 0 && grid[x][y-1][z] == "X" {
		total--
	}
	if y != len(grid[x])-1 && grid[x][y+1][z] == "X" {
		total--
	}
	//check Z
	if z != 0 && grid[x][y][z-1] == "X" {
		total--
	}
	if z != len(grid[x][y])-1 && grid[x][y][z+1] == "X" {
		total--
	}
	return total
}

func isInterior(x, y, z int, grid [20][20][20]string) bool {
	if x == 0 || y == 0 || z == 0 || x == len(grid)-1 || y == len(grid[x])-1 || z == len(grid[x][y])-1 {
		return false
	}

	for i := x; i < len(grid)-1; i++ {
		if grid[i][y][z] == "X" {
			break
		}
	}

	return false
}
