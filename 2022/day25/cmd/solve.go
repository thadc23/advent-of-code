/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func score(scanner *bufio.Scanner, part2 bool) string {

	var nums []int

	for scanner.Scan() {
		line := []rune(scanner.Text())
		nums = append(nums, fromSNAFU(line))
	}

	total := 0

	for _, v := range nums {
		total += v
	}
	return toSNAFU(total)
}

func fromSNAFU(line []rune) int {

	total := 0

	place := int(math.Pow(5, float64(len(line)-1)))

	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '2':
			total += 2 * place
		case '1':
			total += place
		case '-':
			total -= place
		case '=':
			total += -2 * place
		}
		place /= 5
	}
	return total
}

func toSNAFU(num int) string {
	x := num

	var answer string

	var tens string = "0"
	var ones string = "0"

	for x > 0 {
		tens, ones = add(ones, x%5)
		answer = ones + answer
		x /= 5
		ones = tens
		tens = "0"
	}
	if ones != "0" {
		answer = ones + answer
	}
	return answer
}

func add(a string, b int) (string, string) {
	val := lookup(a) + b
	ones := val % 5
	tens := val / 3
	return convert(tens), convert(ones)
}

func lookup(v string) int {
	switch v {
	case "1":
		return 1
	case "2":
		return 2
	case "-":
		return -1
	case "=":
		return -2
	}
	return 0
}
func convert(v int) string {
	switch v {
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "="
	case 4:
		return "-"
	}
	return "0"
}
