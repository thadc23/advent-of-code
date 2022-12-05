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

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Push a new value onto the stack
func (s *Stack) PushAll(str []string) {
	*s = append(*s, str...) // Simply append the new value to the end of the stack
}

// Push a new value onto the stack
func (s *Stack) PushRead(str string) {
	*s = append([]string{str}, *s...) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) PopMultiple(num int) ([]string, bool) {
	if s.IsEmpty() {
		return []string{}, false
	} else {
		index := len(*s) - num   // Get the index of the top most element.
		elements := (*s)[index:] // Index into the slice and obtain the element.
		*s = (*s)[:index]        // Remove it from the stack by slicing it off.
		return elements, true
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
	var stacks []Stack

	for i := 0; i < 9; i++ {
		stacks = append(stacks, Stack{})
	}

	var actions [][]int
	for i := 0; i < 3; i++ {
		actions = append(actions, []int{})
	}

	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum <= 8 {
			for i := 0; i < 9; i++ {
				end := (i * 4) + 4
				if i == 8 {
					end = len(line)
				}
				p := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(line[i*4:end], "[", ""), "]", ""))
				if p != "" {
					stacks[i].PushRead(p)
				}
			}
			lineNum++
		} else if lineNum <= 10 {
			scanner.Text()
			lineNum++
		} else {
			line = strings.ReplaceAll(line, "move", "")
			line = strings.ReplaceAll(line, "from", "")
			line = strings.ReplaceAll(line, "to", "")

			fields := strings.Fields(line)

			numberOfPacakages, _ := strconv.Atoi(fields[0])
			start, _ := strconv.Atoi(fields[1])
			finish, _ := strconv.Atoi(fields[2])
			actions[0] = append(actions[0], numberOfPacakages)
			actions[1] = append(actions[1], start)
			actions[2] = append(actions[2], finish)
		}

	}

	if !part2 {
		for idx := 0; idx < len(actions[0]); idx++ {
			for i := 0; i < actions[0][idx]; i++ {
				packageToMove, exists := stacks[actions[1][idx]-1].Pop()
				if exists {
					stacks[actions[2][idx]-1].PushRead(packageToMove)
				}
			}
		}
	} else {
		for idx := 0; idx < len(actions[0]); idx++ {
			packagesToMove, exists := stacks[actions[1][idx]-1].PopMultiple(actions[0][idx])
			if exists {
				stacks[actions[2][idx]-1].PushAll(packagesToMove)
			}
		}
	}

	final := ""
	for _, s := range stacks {
		val, exists := s.Pop()
		if exists {
			final += val
		}
	}
	return final
}
