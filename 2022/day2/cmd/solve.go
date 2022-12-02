/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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

	fmt.Println(score(scanner, part2))
}

func score(scanner *bufio.Scanner, part2 bool) int {
	var player1 []string
	var player2 []string
	score := 0
	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")
		player1 = append(player1, txt[0])
		player2 = append(player2, txt[1])
	}

	if !part2 {
		for i, v := range player2 {
			switch v {
			case "X":
				switch player1[i] {
				case "A":
					score += 4
				case "B":
					score += 1
				case "C":
					score += 7
				}
			case "Y":
				switch player1[i] {
				case "A":
					score += 8
				case "B":
					score += 5
				case "C":
					score += 2
				}
			case "Z":
				switch player1[i] {
				case "A":
					score += 3
				case "B":
					score += 9
				case "C":
					score += 6
				}
			}
		}
	} else {
		for i, v := range player2 {
			switch v {
			case "X":
				switch player1[i] {
				case "A":
					score += 3
				case "B":
					score += 1
				case "C":
					score += 2
				}
			case "Y":
				switch player1[i] {
				case "A":
					score += 4
				case "B":
					score += 5
				case "C":
					score += 6
				}
			case "Z":
				switch player1[i] {
				case "A":
					score += 8
				case "B":
					score += 9
				case "C":
					score += 7
				}
			}
		}
	}

	return score
}
