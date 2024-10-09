/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

type Galaxy struct {
	num int
	x   float64
	y   float64
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

func calc(scanner *bufio.Scanner, part2 bool) float64 {
	var total float64 = 0
	var grid [][]string

	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	galaxyNum := 0
	var galaxies []Galaxy
	for i, row := range grid {
		for j, col := range row {
			if col == "#" {
				galaxies = append(galaxies, Galaxy{num: galaxyNum, x: float64(j), y: float64(i)})
				galaxyNum++
			}
		}
	}

	expansion := 2
	if part2 {
		expansion = 1000000
	}

	rows, cols := expandGalaxy(grid)

	for i, g := range galaxies {
		if i < len(galaxies) {
			for j := i + 1; j < len(galaxies); j++ {
				total += findDistance(rows, cols, g, galaxies[j], expansion)
			}
		}
	}

	return total
}

func findDistance(rows []int, cols []int, galaxy1, galaxy2 Galaxy, expansion int) float64 {
	var distance float64 = 0

	distance = math.Abs(galaxy1.x-galaxy2.x) + math.Abs(galaxy1.y-galaxy2.y)

	for _, col := range cols {
		if between(float64(col), galaxy1.x, galaxy2.x) {
			distance += float64(expansion - 1)
		}
	}
	for _, row := range rows {
		if between(float64(row), galaxy1.y, galaxy2.y) {
			distance += float64(expansion - 1)
		}
	}

	return distance
}

func between(num, point1, point2 float64) bool {
	if point1 > point2 {
		return num < point1 && num > point2
	}
	return num > point1 && num < point2
}

func expandGalaxy(grid [][]string) ([]int, []int) {

	var rows []int
	var cols []int

	for i := 0; i < len(grid[0]); i++ {
		cols = append(cols, i)
	}

	for i, r := range grid {
		rowEmpty := true
		for j, v := range r {
			if v == "#" {
				rowEmpty = false
				cols = removeElement(cols, j)
			}
		}
		if rowEmpty {
			rows = append(rows, i)
		}
	}

	return rows, cols

}

func removeElement(slice []int, element int) []int {
	var result []int

	for _, value := range slice {
		if value != element {
			result = append(result, value)
		}
	}

	return result
}
