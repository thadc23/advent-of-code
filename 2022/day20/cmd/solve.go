/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

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

	var current []int
	var original []int

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		original = append(original, line)
		current = append(current, line)
	}

	for i := 0; i < len(original); i++ {
		curr := original[i]

		moves := curr % (len(original) - 1)
		if curr == 0 || moves == 0 {
			continue
		}

		index := find(current, curr)
		var newSpot int
		if moves > 0 {
			newSpot = (index + moves) % (len(original) - 1)
		} else if index+moves < 0 {
			newSpot = len(original) + (index + moves) - 1
		} else {
			newSpot = index + moves
		}

		current = remove(current, index)
		current = insert(current, newSpot, curr)
	}

	zeroIndex := find(current, 0)
	oneThousand := (zeroIndex + 1000) % len(current)
	twoThousand := (zeroIndex + 2000) % len(current)
	threeThousand := (zeroIndex + 3000) % len(current)
	fmt.Println(current[oneThousand])
	fmt.Println(current[twoThousand])
	fmt.Println(current[threeThousand])
	return current[oneThousand] + current[twoThousand] + current[threeThousand]
}

func find(current []int, curr int) int {
	for i, v := range current {
		if v == curr {
			return i
		}
	}
	return -1
}

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
