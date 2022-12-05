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
	count := 0
	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")
		elf1 := strings.Split(elves[0], "-")
		elf2 := strings.Split(elves[1], "-")

		left1, _ := strconv.Atoi(elf1[0])
		left2, _ := strconv.Atoi(elf1[1])
		right1, _ := strconv.Atoi(elf2[0])
		right2, _ := strconv.Atoi(elf2[1])

		if !part2 {
			if left1 > right1 && left2 < right2 {
				count++
			} else if left1 == right1 || left2 == right2 {
				count++
			} else if left1 < right1 && left2 > right2 {
				count++
			}
		} else {
			if left1 > right1 && left1 < right2 {
				count++
			} else if left1 == right1 || left2 == right2 || left1 == right2 || left2 == right1 {
				count++
			} else if right1 > left1 && right1 < left2 {
				count++
			}
		}
	}
	return count
}
