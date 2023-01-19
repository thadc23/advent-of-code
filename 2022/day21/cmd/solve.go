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
	"strings"
	"unicode"

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

	monkeys := make(map[string][]string)

	for scanner.Scan() {
		line := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return r == ':' || unicode.IsSpace(r)
		})

		monkeys[line[0]] = line[1:]

	}

	if !part2 {
		return yell(monkeys["root"], monkeys)
	} else {
		x := yell(monkeys[monkeys["root"][0]], monkeys)
		y := yell(monkeys[monkeys["root"][2]], monkeys)

		fmt.Println(x, "==", y)
		return 0
	}
}

func yell(action []string, monkeys map[string][]string) int {
	if len(action) == 1 {
		val, _ := strconv.Atoi(action[0])
		return val
	}

	operand := action[1]

	switch operand {
	case "+":
		return yell(monkeys[action[0]], monkeys) + yell(monkeys[action[2]], monkeys)
	case "-":
		return yell(monkeys[action[0]], monkeys) - yell(monkeys[action[2]], monkeys)
	case "/":
		return yell(monkeys[action[0]], monkeys) / yell(monkeys[action[2]], monkeys)
	case "*":
		return yell(monkeys[action[0]], monkeys) * yell(monkeys[action[2]], monkeys)
	}
	return 0
}

func findOperands(action []string, monkeys map[string][]string, operands [][]string) int {
	if action[0] == "root" {
		return
	}

	

	operand := action[1]
	operands = 

	switch operand {
	case "+":
		return yell(monkeys[action[0]], monkeys) + yell(monkeys[action[2]], monkeys)
	case "-":
		return yell(monkeys[action[0]], monkeys) - yell(monkeys[action[2]], monkeys)
	case "/":
		return yell(monkeys[action[0]], monkeys) / yell(monkeys[action[2]], monkeys)
	case "*":
		return yell(monkeys[action[0]], monkeys) * yell(monkeys[action[2]], monkeys)
	}
	return 0
}
