/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		var input [][]int
		i := 0
		for scanner.Scan() {
			line := scanner.Text()
			input = append(input, make([]int, len(line)))
			for j, x := range line {
				input[i][j], _ = strconv.Atoi(string(x))
			}
			i++
		}

		if part2 {
			solvePart2(input)
		} else {
			solvePart1(input)
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

var min int

func solvePart1(input [][]int) {
	walk(input, len(input)-1, len(input[0])-1, 0)
	fmt.Println(min)
}
func solvePart2(input [][]int) {

}

func walk(input [][]int, x, y, total int) {

	if total+x+y >= min && min > 0 {
		return
	}

	if x == 0 && y == 0 {
		min = total
		// fmt.Println(len(totals))
		return
	}
	total += input[x][y]

	if x > 0 {
		walk(input, x-1, y, total)
	}

	if y > 0 {
		walk(input, x, y-1, total)
	}
}
