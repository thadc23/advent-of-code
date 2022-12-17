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

	var cave = make([][]string, 200)
	maxY := 0
	for i := 0; i < 200; i++ {
		cave[i] = make([]string, 1000)
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")

		for i := 0; i < len(line)-1; i++ {
			start := strings.Split(line[i], ",")
			end := strings.Split(line[i+1], ",")
			startX, _ := strconv.Atoi(start[0])
			startY, _ := strconv.Atoi(start[1])
			endX, _ := strconv.Atoi(end[0])
			endY, _ := strconv.Atoi(end[1])

			if startX == endX {
				var loopStartY int
				var loopEndY int
				if startY > endY {
					loopStartY = endY
					loopEndY = startY
				} else {
					loopStartY = startY
					loopEndY = endY
				}

				if loopEndY > maxY {
					maxY = loopEndY
				}
				//draw vertical line
				for i := loopStartY; i <= loopEndY; i++ {
					cave[i][startX] = "#"
				}
			} else {
				var loopStartX int
				var loopEndX int
				if startX > endX {
					loopStartX = endX
					loopEndX = startX
				} else {
					loopStartX = startX
					loopEndX = endX
				}
				if startY > maxY {
					maxY = startY
				}
				//draw horizontal line
				for i := loopStartX; i <= loopEndX; i++ {
					cave[startY][i] = "#"
				}
			}
		}
	}

	if part2 {
		for i := range cave[maxY+2] {
			cave[maxY+2][i] = "#"
		}
	}

	grains := 0
	x := 500
	y := 0
	for !done(part2, y, maxY, cave) {
		if cave[y+1][x] == "" {
			y++
			continue
		}
		if cave[y+1][x-1] == "" {
			x--
			y++
			continue
		}
		if cave[y+1][x+1] == "" {
			x++
			y++
			continue
		}
		cave[y][x] = "o"
		grains++
		x = 500
		y = 0
	}

	return grains
}

func done(part2 bool, y, maxY int, cave [][]string) bool {
	if part2 {
		return cave[0][500] == "o"
	} else {
		return y >= maxY
	}
}
