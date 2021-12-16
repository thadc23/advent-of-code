/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

		input := make(map[string]string)
		counts := make(map[string]int)

		scanner.Scan()
		polymer := scanner.Text()
		scanner.Scan()

		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " -> ")

			input[parts[0]] = parts[1]

		}

		for _, x := range polymer {
			counts[string(x)]++
		}

		pairs := make(map[string]int)
		for i, x := range polymer {
			if i < len(polymer)-1 {
				pairs[string(x)+string(polymer[i+1])]++
			}
		}

		if part2 {
			solve(pairs, input, counts, 40)
		} else {
			solve(pairs, input, counts, 10)
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

func solve(pairs map[string]int, input map[string]string, counts map[string]int, loops int) {
	newPairs := pairs

	for i := 0; i < loops; i++ {
		curr := newPairs
		newPairs = make(map[string]int)
		for k, v := range curr {
			counts[input[k]] += v
			newPairs[string(k[0])+input[k]] += v
			newPairs[input[k]+string(k[1])] += v
		}
	}

	max := 0
	min := 0
	for _, v := range counts {
		if v > max || max == 0 {
			max = v
		}
		if v < min || min == 0 {
			min = v
		}
	}

	fmt.Println(max - min)
}
