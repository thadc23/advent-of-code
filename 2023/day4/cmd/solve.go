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

	fmt.Println(calc(scanner, part2))
}

func calc(scanner *bufio.Scanner, part2 bool) int {
	var total_tickets map[int]int = make(map[int]int)

	for i := 1; i <= 209; i++ {
		total_tickets[i] = 1
	}
	var total int = 0
	i := 1
	for scanner.Scan() {
		input := scanner.Text()

		// Find the index of the colon
		colonIndex := strings.Index(input, ":")

		// Extract the substring after the colon
		substring := input[colonIndex+1:]

		// Split the remaining substring based on the vertical bar (|)
		parts := strings.Split(substring, "|")

		// Parse the first slice of integers
		firstSlice, _ := parseIntSlice(parts[0])

		// Parse the second slice of integers
		secondSlice, _ := parseIntSlice(parts[1])

		// Find and print the number of matches between the two slices
		matches := findMatches(firstSlice, secondSlice)
		if !part2 && matches > 0 {
			total += int(math.Pow(2, float64(matches-1)))
		} else if part2 {
			for y := 0; y < total_tickets[i]; y++ {
				for x := 1; x <= matches; x++ {
					total_tickets[i+x]++
				}
			}
		}
		i++
	}

	if part2 {
		return calculateTotalTickets(total_tickets)
	}
	return total
}

func calculateTotalTickets(total_tickets map[int]int) int {
	total := 0
	for _, v := range total_tickets {
		total += v
	}
	return total
}

// parseIntSlice parses a space-separated string of integers into a slice of integers
func parseIntSlice(s string) ([]int, error) {
	fields := strings.Fields(s)
	intSlice := make([]int, len(fields))

	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, fmt.Errorf("error converting string to int: %v", err)
		}
		intSlice[i] = num
	}

	return intSlice, nil
}

// findMatches finds the number of matches between two slices of integers
func findMatches(slice1, slice2 []int) int {
	matchCount := 0

	// Create a map to store the occurrences of each number in the first slice
	occurrences := make(map[int]int)
	for _, num := range slice1 {
		occurrences[num]++
	}

	// Check the occurrences in the second slice
	for _, num := range slice2 {
		if count, ok := occurrences[num]; ok && count > 0 {
			matchCount++
			occurrences[num]--
		}
	}

	return matchCount
}
