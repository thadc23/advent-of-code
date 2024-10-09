/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

	fmt.Println(calc(scanner, part2))
}

var digit_words = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func calc(scanner *bufio.Scanner, part2 bool) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := 0
		first_idx := len(line)
		last := 0
		last_idx := -1

		if part2 {
			digits = append(digit_words, digits...)
		}

		for i, v := range digits {
			idx := strings.Index(line, v)
			if idx == -1 {
				continue
			}
			if idx < first_idx {
				first_idx = idx
				first = i % 10
			}

			idx = strings.LastIndex(line, v)
			if idx > last_idx {
				last_idx = idx
				last = i % 10
			}
		}
		total += (first * 10) + last
	}
	return total
}
