/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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

type Fold struct {
	axis string
	line int
}

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

		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		input := make(map[int]map[int]string)
		var folds []Fold

		for scanner.Scan() {
			line := scanner.Text()

			parts := strings.Split(line, ",")

			if len(parts) == 2 {
				x, _ := strconv.Atoi(parts[0])
				y, _ := strconv.Atoi(parts[1])

				if _, ok := input[x]; !ok {
					input[x] = make(map[int]string)
				}
				input[x][y] = "#"
			} else if len(line) != 0 {
				parts = strings.Split(line, " ")
				fold := strings.Split(parts[2], "=")
				foldLine, _ := strconv.Atoi(fold[1])
				folds = append(folds, Fold{axis: fold[0], line: foldLine})
			}
		}

		if part2 {
			solvePart2(input, folds)
		} else {
			solvePart1(input, folds)
		}
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	solveCmd.PersistentFlags().Bool("part2", false, "Solve part2")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func solvePart1(input map[int]map[int]string, folds []Fold) {
	input = foldPaper(input, folds[0])

	count := 0
	for i, x := range input {
		for j, _ := range x {
			if input[i][j] == "#" {
				count++
			}
		}
	}

	fmt.Println(count)
}
func solvePart2(input map[int]map[int]string, folds []Fold) {

	paper := input
	for _, f := range folds {
		paper = foldPaper(paper, f)
	}

	maxX, maxY := findMax(paper)

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if paper[x][y] != "#" {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}

}

func foldPaper(input map[int]map[int]string, fold Fold) map[int]map[int]string {
	switch fold.axis {
	case "x":
		for i, x := range input {
			for j, _ := range x {
				if i > fold.line && input[i][j] == "#" {
					if _, ok := input[fold.line-(i-fold.line)]; !ok {
						input[fold.line-(i-fold.line)] = make(map[int]string)
					}
					input[fold.line-(i-fold.line)][j] = "#"
					input[i][j] = " "
				}
			}
		}
	case "y":
		for i, x := range input {
			for j, _ := range x {
				if j > fold.line && input[i][j] == "#" {
					input[i][fold.line-(j-fold.line)] = "#"
					input[i][j] = " "
				}
			}
		}
	}
	return input
}

func findMax(input map[int]map[int]string) (int, int) {

	maxX := 0
	maxY := 0

	for x, _ := range input {
		for y, _ := range input[x] {
			if input[x][y] == "#" {
				if y > maxY {
					maxY = y
				}
				if x > maxX {
					maxX = x
				}
			}
		}
	}

	return maxX, maxY
}
