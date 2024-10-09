/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

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

type Node struct {
	name  string
	left  string
	right string
}

func calc(scanner *bufio.Scanner, part2 bool) int {

	var nodes map[string]Node = make(map[string]Node)
	var start []Node

	scanner.Scan()
	moves := strings.Split(strings.TrimSpace(scanner.Text()), "")
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		node_name := strings.Split(line, " = ")[0]
		left := strings.Trim(strings.Split(strings.Split(line, " = ")[1], ",")[0], " ()")
		right := strings.Trim(strings.Split(strings.Split(line, " = ")[1], ",")[1], " ()")
		node := Node{name: node_name, left: left, right: right}
		if part2 && strings.Split(node_name, "")[2] == "A" {
			start = append(start, node)
		}
		nodes[node_name] = node
	}

	if part2 {
		return findSteps(nodes, moves, start)
	}
	return findSteps(nodes, moves, []Node{nodes["AAA"]})
}

func findSteps(nodes map[string]Node, moves []string, start []Node) int {
	var wg sync.WaitGroup
	resultChan := make(chan int, len(start))

	// Launch goroutines
	for i := 0; i < len(start); i++ {
		wg.Add(1)
		go walk(nodes, moves, start[i], &wg, resultChan)
	}

	// Close the result channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var step_counts []int

	for n := range resultChan {
		step_counts = append(step_counts, n)
	}

	return findLCM(step_counts)
}

func walk(nodes map[string]Node, moves []string, start Node, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	curr := start
	steps := 0
	for i := 0; !isEnd(curr); i = (i + 1) % len(moves) {
		steps++
		switch moves[i] {
		case "L":
			curr = nodes[curr.left]
		case "R":
			curr = nodes[curr.right]
		}
	}

	resultChan <- steps
}

func isEnd(node Node) bool {
	return node.name[len(node.name)-1:] == "Z"
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find the Least Common Multiple (LCM) using GCD
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to find the LCM of multiple numbers
func findLCM(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result = lcm(result, num)
	}
	return result
}
