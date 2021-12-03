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

		solve(part2, scanner)
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

func solve(part2 bool, scanner *bufio.Scanner) {

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if part2 {
		solvePart2(input)
	} else {
		solvePart1(input)
	}

}

func solvePart1(input []string) {
	ones, zeros := findOnesAndZeros(input)

	gamma := 0
	epsilon := 0
	i := 0
	for i < 12 {
		if ones[i] > zeros[i] {
			gamma = (gamma << 1) + 1
			epsilon = epsilon << 1
		} else {
			gamma = gamma << 1
			epsilon = (epsilon << 1) + 1
		}
		i++
	}

	fmt.Println(gamma * epsilon)
}

func solvePart2(input []string) {

	var oxygen, co2 []string
	for _, val := range input {
		oxygen = append(oxygen, val)
		co2 = append(co2, val)
	}

	i := 0
	for len(oxygen) > 1 {
		ones, zeros := findOnesAndZeros(oxygen)
		if ones[i] >= zeros[i] {
			oxygen = filter(oxygen, i, '1')
		} else {
			oxygen = filter(oxygen, i, '0')
		}
		i++
	}

	i = 0
	for len(co2) > 1 {
		ones, zeros := findOnesAndZeros(co2)
		if ones[i] >= zeros[i] {
			co2 = filter(co2, i, '0')
		} else {
			co2 = filter(co2, i, '1')
		}
		i++
	}

	fmt.Println(oxygen)
	fmt.Println(co2)
	oxygenInt, _ := strconv.ParseInt(oxygen[0], 2, 64)
	co2Int, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Println(oxygenInt * co2Int)
}

func filter(values []string, position int, val rune) []string {
	if len(values) == 1 {
		return values
	}

	var retVal []string
	for _, input := range values {
		if []rune(input)[position] == val {
			retVal = append(retVal, input)
		}
	}

	return retVal
}

func findOnesAndZeros(input []string) ([]int, []int) {
	ones := make([]int, 12)
	zeros := make([]int, 12)

	for _, val := range input {
		current, _ := strconv.ParseInt(val, 2, 64)
		i := 0
		for i < 12 {
			bit := current % 2
			current = current >> 1
			if bit == 1 {
				ones[11-i] = ones[11-i] + 1
			} else {
				zeros[11-i] = zeros[11-i] + 1
			}
			i++
		}
	}
	return ones, zeros
}
