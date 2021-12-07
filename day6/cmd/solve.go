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

		fish := make(map[string]int)
		for scanner.Scan() {
			slice := strings.Split(scanner.Text(), ",")
			for _, age := range slice {
				fish[age] += 1
			}
		}

		if part2 {
			solvePart2(fish)
		} else {
			solvePart1(fish)
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

func solvePart1(fish map[string]int) {
	loop(fish, 80)
}
func solvePart2(fish map[string]int) {
	loop(fish, 256)
}

func loop(fish map[string]int, times int) {

	for i := 0; i < times; i++ {
		for j := 0; j < 10; j++ {
			if j == 0 {
				fish["7"] += fish["0"]
				fish["9"] += fish["0"]
				fish["0"] = 0
				continue
			}
			fish[fmt.Sprint(j-1)] += fish[fmt.Sprint(j)]
			fish[fmt.Sprint(j)] = 0
		}
	}
	sum := 0
	for j := 0; j < 10; j++ {
		sum += fish[fmt.Sprint(j)]
	}
	fmt.Println(sum)

}
