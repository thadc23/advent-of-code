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

	var commands []int
	var grid [][]string

	if part2 {
		for i := 0; i < 6; i++ {
			grid = append(grid, []string{})
			for j := 0; j < 40; j++ {
				grid[i] = append(grid[i], ".")
			}
		}
	}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) > 1 {
			val, _ := strconv.Atoi(line[1])
			commands = append(commands, val)
		} else {
			commands = append(commands, 0)
		}
	}

	sum := 0
	X := 1
	cycles := 1

	for i := 0; i < len(commands); i++ {
		if (cycles-20)%40 == 0 {
			sum += X * cycles
		}
		gridY := (cycles - 1) / 40
		gridX := (cycles - 1) % 40

		if gridX == X || gridX == X-1 || gridX == X+1 {
			grid[gridY][gridX] = "#"
		}
		if commands[i] != 0 {
			cycles++
			if (cycles-20)%40 == 0 {
				sum += X * cycles
			}
			gridY := (cycles - 1) / 40
			gridX := (cycles - 1) % 40

			if gridX == X || gridX == X-1 || gridX == X+1 {
				grid[gridY][gridX] = "#"
			}
		}
		X += commands[i]
		cycles++
	}
	fmt.Println(grid)
	return sum
}
