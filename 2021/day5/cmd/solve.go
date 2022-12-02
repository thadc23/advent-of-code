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

		if part2 {
			solvePart2(scanner)
		} else {
			solvePart1(scanner)
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

func solvePart1(scanner *bufio.Scanner) {

	board := make([][]int, 1000)

	for i := 0; i < 1000; i++ {
		board[i] = make([]int, 1000)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		startPoint := parts[0]
		endPoint := parts[1]

		startXStr := strings.Split(startPoint, ",")[0]
		startYStr := strings.Split(startPoint, ",")[1]
		endXStr := strings.Split(endPoint, ",")[0]
		endYStr := strings.Split(endPoint, ",")[1]

		startXInt, _ := strconv.Atoi(startXStr)
		startYInt, _ := strconv.Atoi(startYStr)
		endXInt, _ := strconv.Atoi(endXStr)
		endYInt, _ := strconv.Atoi(endYStr)

		if startXInt == endXInt {
			startY := startYInt
			endY := endYInt
			if startYInt > endYInt {
				startY = endYInt
				endY = startYInt
			}
			for i := startY; i <= endY; i++ {
				board[startXInt][i] = board[startXInt][i] + 1
			}
		} else if startYInt == endYInt {
			startX := startXInt
			endX := endXInt
			if startXInt > endXInt {
				startX = endXInt
				endX = startXInt
			}
			for i := startX; i <= endX; i++ {
				board[i][startYInt] = board[i][startYInt] + 1
			}
		}

	}
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if board[i][j] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)

}
func solvePart2(scanner *bufio.Scanner) {
	board := make([][]int, 1000)

	for i := 0; i < 1000; i++ {
		board[i] = make([]int, 1000)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		startPoint := parts[0]
		endPoint := parts[1]

		startXStr := strings.Split(startPoint, ",")[0]
		startYStr := strings.Split(startPoint, ",")[1]
		endXStr := strings.Split(endPoint, ",")[0]
		endYStr := strings.Split(endPoint, ",")[1]

		startXInt, _ := strconv.Atoi(startXStr)
		startYInt, _ := strconv.Atoi(startYStr)
		endXInt, _ := strconv.Atoi(endXStr)
		endYInt, _ := strconv.Atoi(endYStr)

		if startXInt == endXInt {
			startY := startYInt
			endY := endYInt
			if startYInt > endYInt {
				startY = endYInt
				endY = startYInt
			}
			for i := startY; i <= endY; i++ {
				board[startXInt][i] = board[startXInt][i] + 1
			}
		} else if startYInt == endYInt {
			startX := startXInt
			endX := endXInt
			if startXInt > endXInt {
				startX = endXInt
				endX = startXInt
			}
			for i := startX; i <= endX; i++ {
				board[i][startYInt] = board[i][startYInt] + 1
			}
		} else {
			startY := startYInt
			startX := startXInt
			endX := endXInt
			endY := endYInt
			if startXInt > endXInt {
				startX = endXInt
				endX = startXInt
				startY = endYInt
				endY = startYInt
			}
			for startX <= endX {
				board[startX][startY] = board[startX][startY] + 1
				if startY > endY {
					startY--
				} else {
					startY++
				}
				startX++
			}

		}

	}
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if board[i][j] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
