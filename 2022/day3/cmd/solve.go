/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/juliangruber/go-intersect/v2"
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
	var priority = 0
	if !part2 {
		var sacks1 [][]rune
		var sacks2 [][]rune
		for scanner.Scan() {
			txt := scanner.Text()
			runes := []rune(txt)
			sacks1 = append(sacks1, runes[0:len(runes)/2])
			sacks2 = append(sacks2, runes[len(runes)/2:])
		}

		for i := range sacks1 {
			val := int(intersect.SimpleGeneric(sacks1[i], sacks2[i])[0])
			priority += calculatePriority(val)
		}
	} else {
		var sacks [][]rune
		for scanner.Scan() {
			sacks = append(sacks, []rune(scanner.Text()))
		}

		for i := 2; i < len(sacks); i += 3 {
			val := int(intersect.SimpleGeneric(intersect.SimpleGeneric(sacks[i], sacks[i-1]), sacks[i-2])[0])
			priority += calculatePriority(val)
		}
	}

	return priority
}

func calculatePriority(val int) int {
	if val >= 97 {
		return val - 'a' + 1
	}
	return val - 'A' + 27
}
