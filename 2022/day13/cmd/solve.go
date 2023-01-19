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

type Node struct {
	Val   int
	Slice []Node
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

func score(scanner *bufio.Scanner, part2 bool) int {

	for i := 1; scanner.Scan(); i++ {
		left := scanner.Text()
		scanner.Scan()
		right := scanner.Text()
		scanner.Scan()
	}
	return 0
}

func createNodes(line string) Node {
	var nodes []Node
	var root = Node{}
	var currNode Node
	var currSlice []Node
	for _, c := range strings.TrimSuffix(strings.TrimPrefix(line, "["), "]") {
		switch c {
		case '[':
			currSlice = nodes[curr].Slice
		case ',':
			currNode
			nodes = append(nodes, Node{})
			curr++
		case ']':
			currSlice = nodes[curr].Slice
		default:
			val, _ := strconv.Atoi(string(c))
		}
	}
	return root
}

func testInts(left, right int) int {
	return right - left
}

func testIntSlices(left, right []int) int {
	for i, v := range left {
		if len(right)-1 < i {
			return 1
		}
		result := testInts(v, right[i])
		if result > 0 {
			return 1
		} else if result < 0 {
			return -1
		}
	}
}
