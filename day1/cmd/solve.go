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
		slidingScale, _ := cmd.Flags().GetBool("sliding")
		solve(slidingScale)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	solveCmd.PersistentFlags().Bool("sliding", false, "Solve using a sliding window")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func solve(slidingScale bool) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if slidingScale {
		fmt.Println(solveSliding(scanner))
	} else {
		fmt.Println(solveStandard(scanner))
	}
}

func solveSliding(scanner *bufio.Scanner) int {
	var readings []int

	for scanner.Scan() {
		cur, _ := strconv.Atoi(scanner.Text())
		readings = append(readings, cur)
	}
	var groupings []int
	for i, _ := range readings {
		if i+1 < len(readings) && i+2 < len(readings) {
			total := readings[i] + readings[i+1] + readings[i+2]
			groupings = append(groupings, total)
		}
	}

	increases := 0
	prev := -1
	for _, cur := range groupings {
		if cur > prev && prev != -1 {
			increases++
		}
		prev = cur
	}
	return increases
}

func solveStandard(scanner *bufio.Scanner) int {
	increases := 0
	prev := -1
	for scanner.Scan() {
		cur, _ := strconv.Atoi(scanner.Text())
		if cur > prev && prev != -1 {
			increases++
		}
		prev = cur
	}
	return increases
}
