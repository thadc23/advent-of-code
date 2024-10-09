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

var max = map[string]int{"red": 12, "green": 13, "blue": 14}

func calc(scanner *bufio.Scanner, part2 bool) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_slice := strings.Split(line, ":")
		game_id, _ := strconv.Atoi(strings.Split(line_slice[0], " ")[1])
		rounds := strings.Split(line_slice[1], ";")
		if !part2 {
			if run(rounds) {
				total += game_id
			}
		} else {
			total += findPower(rounds)
		}
	}

	return total
}

func run(rounds []string) bool {
	for _, r := range rounds {
		colors := strings.Split(r, ",")
		for _, c := range colors {
			color_slice := strings.Split(strings.TrimSpace(c), " ")
			count, _ := strconv.Atoi(color_slice[0])
			if max[color_slice[1]] < count {
				return false
			}
		}
	}
	return true
}

func findPower(rounds []string) int {
	var highest = map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, r := range rounds {
		colors := strings.Split(r, ",")
		for _, c := range colors {
			color_slice := strings.Split(strings.TrimSpace(c), " ")
			count, _ := strconv.Atoi(color_slice[0])

			if highest[color_slice[1]] < count {
				highest[color_slice[1]] = count
			}
		}

	}
	return highest["red"] * highest["green"] * highest["blue"]
}
