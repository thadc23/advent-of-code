/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

type Elf struct {
	cals int
}

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

	fmt.Println(mostCalories(scanner, part2))
}

func mostCalories(scanner *bufio.Scanner, part2 bool) int {
	var elves []Elf
	current_elf := Elf{cals: 0}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt != "" {
			cur, _ := strconv.Atoi(txt)
			current_elf.cals += cur
		} else {
			elves = append(elves, current_elf)
			current_elf = Elf{cals: 0}
		}
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].cals > elves[j].cals
	})
	if part2 {
		return elves[0].cals + elves[1].cals + elves[2].cals
	}
	return elves[0].cals
}
