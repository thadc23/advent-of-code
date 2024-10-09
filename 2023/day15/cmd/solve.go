/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

	fmt.Println(calc(scanner, part2))
}

func calc(scanner *bufio.Scanner, part2 bool) int {
	total := 0
	var items []string
	for scanner.Scan() {
		items = strings.Split(scanner.Text(), ",")
	}

	if part2 {
		return solve_part2(items)
	}

	for _, i := range items {
		current := 0
		for _, char := range i {
			current += int(char)
			current *= 17
			current = current % 256
		}
		total += current
	}
	return total
}

type Lens struct {
	name         string
	focal_length int
}

func solve_part2(items []string) int {
	total := 0
	var boxes [][]Lens

	for i := 0; i < 256; i++ {
		boxes = append(boxes, make([]Lens, 0))
	}

	for _, i := range items {
		eq := strings.Split(i, "=")
		lens_name := strings.TrimRight(eq[0], "-")
		box_num := hash(lens_name)
		index := findLens(boxes[box_num], lens_name)

		if len(eq) == 2 {
			focal_length, _ := strconv.Atoi(eq[1])
			if index == -1 {
				boxes[box_num] = append(boxes[box_num], Lens{name: lens_name, focal_length: focal_length})
			} else {
				boxes[box_num][index] = Lens{name: lens_name, focal_length: focal_length}
			}
		} else {
			if index != -1 {
				copy(boxes[box_num][index:], boxes[box_num][index+1:])
				boxes[box_num] = boxes[box_num][:len(boxes[box_num])-1]
			}
		}
	}

	for box_num, box := range boxes {
		for slot, lens := range box {
			total += (1 + box_num) * (1 + slot) * lens.focal_length
		}
	}

	return total
}

func findLens(lenses []Lens, name string) int {
	for i, l := range lenses {
		if l.name == name {
			return i
		}
	}
	return -1
}

func hash(value string) int {
	total := 0
	for _, char := range value {
		total += int(char)
		total *= 17
		total = total % 256
	}
	return total
}
