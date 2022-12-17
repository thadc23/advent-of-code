/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
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

func score(scanner *bufio.Scanner, part2 bool) int64 {

	var grid [][]rune
	startNode := 0
	endNode := 0

	nodeNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	graphString := ""

	for y, row := range grid {
		for x, col := range row {
			if col == rune('S') {
				startNode = nodeNum
				grid[y][x] = rune('a')
			} else if col == rune('E') {
				endNode = nodeNum
				grid[y][x] = rune('z')
			}
			nodeNum++
		}
	}

	nodeNum = 0
	for y, row := range grid {
		for x, col := range row {
			graphString += strconv.Itoa(nodeNum) + " "
			if y > 0 && grid[y-1][x] <= col+1 {
				graphString += strconv.Itoa(nodeNum-len(grid[0])) + ",1 "
			}
			if y < len(grid)-1 && grid[y+1][x] <= col+1 {
				graphString += strconv.Itoa(nodeNum+len(grid[0])) + ",1 "
			}
			if x > 0 && grid[y][x-1] <= col+1 {
				graphString += strconv.Itoa(nodeNum-1) + ",1 "
			}
			if x < len(grid[y])-1 && grid[y][x+1] <= col+1 {
				graphString += strconv.Itoa(nodeNum+1) + ",1 "
			}
			graphString += "\n"
			nodeNum++
		}
	}
	os.WriteFile("./graph.txt", []byte(graphString), 0644)
	graph, _ := dijkstra.Import("./graph.txt")
	if part2 {
		var paths []int64
		nodeNum = 0
		for _, row := range grid {
			for _, col := range row {
				if col == rune('a') {
					min, err := graph.Shortest(nodeNum, endNode)
					if err == nil {
						paths = append(paths, min.Distance)
					}
				}
				nodeNum++
			}
		}

		sort.Slice(paths, func(i, j int) bool {
			return paths[i] < paths[j]
		})
		return paths[0]
	} else {
		min, _ := graph.Shortest(startNode, endNode)
		return min.Distance
	}
}
