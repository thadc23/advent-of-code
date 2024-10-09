/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

	fmt.Println(calc(scanner, part2))
}

func calc(scanner *bufio.Scanner, part2 bool) int {
	total := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		total += findNext(convertStringsToInts(line), part2)
	}

	return total
}

func findNext(s []int, part2 bool) int {
	var levels [][]int = make([][]int, 0)
	levels = append(levels, s)
	curr := s
	for notAllZeroes(curr) {
		curr = buildNextLevel(curr)
		levels = append(levels, curr)
	}

	if !part2 {
		for i := len(levels) - 1; i > 0; i-- {
			levels[i-1] = append(levels[i-1], levels[i][len(levels[i])-1]+levels[i-1][len(levels[i-1])-1])
		}
		return levels[0][len(levels[0])-1]
	} else {
		for i := len(levels) - 1; i > 0; i-- {
			levels[i-1] = prependElement(levels[i-1], levels[i-1][0]-levels[i][0])
		}
		return levels[0][0]
	}

}

func prependElement(slice []int, element int) []int {
	// Create a new slice with the element at the beginning
	newSlice := make([]int, 1)
	newSlice[0] = element

	// Append the existing slice to the new slice
	newSlice = append(newSlice, slice...)

	return newSlice
}

func notAllZeroes(s []int) bool {
	for _, v := range s {
		if v != 0 {
			return true
		}
	}
	return false
}

func buildNextLevel(s []int) []int {
	var nextLevel []int
	for i := 0; i < len(s)-1; i++ {
		nextLevel = append(nextLevel, s[i+1]-s[i])
	}
	return nextLevel

}

func convertStringsToInts(strArray []string) []int {
	var intArray []int

	for _, str := range strArray {
		// Parse string to int
		num, _ := strconv.Atoi(str)
		// Append the parsed int to the new array
		intArray = append(intArray, num)
	}

	return intArray
}
