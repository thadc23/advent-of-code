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
	"sort"
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

		var crabs []int
		for scanner.Scan() {
			slice := strings.Split(scanner.Text(), ",")
			for _, position := range slice {
				i, _ := strconv.Atoi(position)
				crabs = append(crabs, i)
			}
		}

		if part2 {
			solvePart2(crabs)
		} else {
			solvePart1(crabs)
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

//Works, but should it?
func solvePart1(crabs []int) {
	sort.Ints(crabs)
	median := crabs[len(crabs)/2]
	fuel := 0
	for _, p := range crabs {
		if p > median {
			fuel += p - median
		} else if p < median {
			fuel += median - p
		}
	}
	fmt.Println(fuel)
}

func solvePart2(crabs []int) {
	max, min := findMaxAndMin(crabs)

	fuelMap := make(map[int]int)

	for i := min; i <= max; i++ {
		fuelMap[i] = 0
	}

	for i := min; i <= max; i++ {
		for _, p := range crabs {
			if p > i {
				fuelMap[i] += sumSequence(p - i)
			} else if p < i {
				fuelMap[i] += sumSequence(i - p)
			}
		}
	}
	result := fuelMap[min]
	for i := min; i <= max; i++ {
		if fuelMap[i] < result {
			result = fuelMap[i]
		}
	}
	fmt.Println(result)

}

func sumSequence(n int) int {
	return (n * (n + 1)) / 2
}

func findMaxAndMin(crabs []int) (int, int) {
	min := 0
	max := 0
	for _, p := range crabs {
		if p > max {
			max = p
		}
		if p < min {
			min = p
		}
	}
	return max, min
}
