/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(r rune) {
	*s = append(*s, r) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
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

		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		var input []string

		for scanner.Scan() {
			input = append(input, scanner.Text())
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

func solvePart1(input []string) {
	total := 0

	points := make(map[rune]int)
	points[')'] = 3
	points[']'] = 57
	points['}'] = 1197
	points['>'] = 25137

	for _, line := range input {
		var stack Stack
		for _, r := range line {
			switch r {
			case '{':
				stack.Push('}')
			case '[':
				stack.Push(']')
			case '<':
				stack.Push('>')
			case '(':
				stack.Push(')')
			case '}', ']', '>', ')':
				expected, _ := stack.Pop()
				if expected != r {
					total += points[r]
					break
				}
			}
		}
	}

	fmt.Println(total)
}
func solvePart2(input []string) {
	var total []int

	points := make(map[rune]int)
	points[')'] = 1
	points[']'] = 2
	points['}'] = 3
	points['>'] = 4

	for _, line := range input {
		corrupt := false
		var stack Stack
		for _, r := range line {
			switch r {
			case '{':
				stack.Push('}')
			case '[':
				stack.Push(']')
			case '<':
				stack.Push('>')
			case '(':
				stack.Push(')')
			case '}', ']', '>', ')':
				expected, _ := stack.Pop()
				if expected != r {
					corrupt = true
				}
			}
		}

		if !corrupt {
			score := 0
			for true {
				a, more := stack.Pop()
				if !more {
					break
				}
				score = ((score * 5) + points[a])
			}
			total = append(total, score)
		}

	}

	sort.Ints(total)
	fmt.Println(total[len(total)/2])
}
