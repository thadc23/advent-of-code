/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

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

	fmt.Println(score(scanner, part2))
}

func score(scanner *bufio.Scanner, part2 bool) int {

	var chars []rune
	for scanner.Scan() {
		chars = []rune(scanner.Text())
	}

	if part2 {
		return findDistinct(chars, 14)
	} else {
		return findDistinct(chars, 4)
	}
}

func findDistinct(chars []rune, size int) int {
	var buffer []rune
	for i, c := range chars {
		buffer = append(buffer, c)
		if len(buffer) == size {
			if duplicateInArray(buffer) == -1 {
				return i + 1
			} else {
				buffer = buffer[1:]
			}
		}
	}
	return -1
}

func duplicateInArray(arr []rune) rune {
	visited := make(map[rune]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] {
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return -1
}
