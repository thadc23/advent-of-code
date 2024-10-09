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
	total := 1

	scanner.Scan()
	times := convertStringsToInts(strings.Fields(strings.Split(scanner.Text(), ":")[1]))
	scanner.Scan()
	distances := convertStringsToInts(strings.Fields(strings.Split(scanner.Text(), ":")[1]))

	if part2 {
		times = combine(times)
		distances = combine(distances)
	}

	fmt.Println(times)
	fmt.Println(distances)

	for i := 0; i < len(times); i++ {
		wins := findWins(times[i], distances[i])
		total *= wins
	}

	return total
}

func combine(slice []int) []int {
	total := ""

	for _, v := range slice {
		total += strconv.Itoa(v)
	}

	total_int, _ := strconv.Atoi(total)
	return []int{total_int}

}

func findWins(time, record int) int {
	wins := 0
	for j := 0; j < time; j++ {
		distance := (time - j) * j
		if distance > record {
			wins++
		}
	}
	return wins
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
