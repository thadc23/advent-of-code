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
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

const (
	MULTIPLY int = 0
	ADD      int = 1
	SQUARE   int = 2
)

type Monkey struct {
	Items     []int
	Op        int
	Val       int
	Divisor   int
	Yes       int
	No        int
	Inspected int
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
	var monkeys []*Monkey
	var rounds int
	if part2 {
		rounds = 10000
	} else {
		rounds = 20
	}

	for scanner.Scan() {
		scanner.Scan()
		m := Monkey{}
		line := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return unicode.IsSpace(r) || r == ','
		})
		items := make([]int, len(line[2:]))

		for i, s := range line[2:] {
			val, _ := strconv.Atoi(s)
			items[i] = val
		}
		m.Items = items

		scanner.Scan()
		line = strings.Fields(scanner.Text())
		val, err := strconv.Atoi(line[len(line)-1])
		switch line[len(line)-2] {
		case "*":
			if err != nil {
				m.Op = SQUARE
			} else {
				m.Op = MULTIPLY
				m.Val = val
			}
		case "+":
			m.Op = ADD
			m.Val = val
		}

		scanner.Scan()
		line = strings.Fields(scanner.Text())
		val, _ = strconv.Atoi(line[len(line)-1])
		m.Divisor = val

		scanner.Scan()
		line = strings.Fields(scanner.Text())
		yes, _ := strconv.Atoi(line[len(line)-1])
		m.Yes = yes

		scanner.Scan()
		line = strings.Fields(scanner.Text())
		no, _ := strconv.Atoi(line[len(line)-1])
		m.No = no

		scanner.Scan()

		monkeys = append(monkeys, &m)
	}

	for round := 0; round < rounds; round++ {
		commonMultiple := getCommonMultiple(monkeys)
		for _, m := range monkeys {
			for _, item := range m.Items {
				m.Inspected++
				switch m.Op {
				case ADD:
					item += m.Val
				case MULTIPLY:
					item *= m.Val
				case SQUARE:
					item *= item
				}
				if !part2 {
					item /= 3
				} else {
					item %= commonMultiple
				}
				if item%m.Divisor == 0 {
					monkeys[m.Yes].Items = append(monkeys[m.Yes].Items, item)
				} else {
					monkeys[m.No].Items = append(monkeys[m.No].Items, item)
				}
			}
			m.Items = []int{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})

	return monkeys[0].Inspected * monkeys[1].Inspected
}

func getCommonMultiple(monkeys []*Monkey) int {
	common_multiple := 1
	for _, m := range monkeys {
		common_multiple *= m.Divisor
	}

	return common_multiple
}
