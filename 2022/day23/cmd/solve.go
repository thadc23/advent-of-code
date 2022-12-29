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

type Position struct {
	Row int
	Col int
}

type Direction struct {
	Name string
}

type Elf struct {
	Current Position
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

	directions := []Direction{
		Direction{Name: "North"},
		Direction{Name: "South"},
		Direction{Name: "West"},
		Direction{Name: "East"},
	}

	row := 0

	var elves []*Elf

	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			if v == '#' {
				elves = append(elves, &Elf{Current: Position{Row: row, Col: i}})
			}
		}
		row++
	}

	var currentLocations map[string]string
	var nextLocations map[string][]*Elf

	rounds := 1
	for {
		currentLocations = make(map[string]string)
		nextLocations = make(map[string][]*Elf)
		for _, elf := range elves {
			currentLocations[fmt.Sprintf("%d,%d", elf.Current.Row, elf.Current.Col)] = "x"
		}
		for _, elf := range elves {
			move(directions, elf, currentLocations, nextLocations)
		}

		didMove := false
		for k, v := range nextLocations {
			if len(v) == 1 {
				row, _ := strconv.Atoi(strings.Split(k, ",")[0])
				col, _ := strconv.Atoi(strings.Split(k, ",")[1])
				v[0].Current.Row = row
				v[0].Current.Col = col
				didMove = true
			}
		}

		if part2 && !didMove {
			return rounds
		}
		if !part2 && rounds == 10 {
			break
		}
		rounds++
		directions = append(directions[1:], directions[0])
	}

	minRow := 100000
	maxRow := -100000
	minCol := 100000
	maxCol := -100000

	currentLocations = make(map[string]string)
	for _, elf := range elves {
		currentLocations[fmt.Sprintf("%d,%d", elf.Current.Row, elf.Current.Col)] = "x"
	}

	for _, elf := range elves {
		if elf.Current.Row < minRow {
			minRow = elf.Current.Row
		}
		if elf.Current.Row > maxRow {
			maxRow = elf.Current.Row
		}
		if elf.Current.Col < minCol {
			minCol = elf.Current.Col
		}
		if elf.Current.Col > maxCol {
			maxCol = elf.Current.Col
		}
	}

	printGrid(minRow, maxRow, minCol, maxCol, currentLocations)

	return (maxRow-minRow+1)*(maxCol-minCol+1) - len(elves)
}

func move(directions []Direction, elf *Elf, currentLocations map[string]string, nextLocations map[string][]*Elf) {
	for _, d := range directions {
		if !elf.isBlocked(currentLocations, d) {
			elf.simulateMove(d, nextLocations)
			return
		}
	}
}

func (elf *Elf) isBlocked(currentLocations map[string]string, d Direction) bool {

	nw := fmt.Sprintf("%d,%d", elf.Current.Row-1, elf.Current.Col-1)
	n := fmt.Sprintf("%d,%d", elf.Current.Row-1, elf.Current.Col)
	ne := fmt.Sprintf("%d,%d", elf.Current.Row-1, elf.Current.Col+1)
	e := fmt.Sprintf("%d,%d", elf.Current.Row, elf.Current.Col+1)
	se := fmt.Sprintf("%d,%d", elf.Current.Row+1, elf.Current.Col+1)
	s := fmt.Sprintf("%d,%d", elf.Current.Row+1, elf.Current.Col)
	sw := fmt.Sprintf("%d,%d", elf.Current.Row+1, elf.Current.Col-1)
	w := fmt.Sprintf("%d,%d", elf.Current.Row, elf.Current.Col-1)

	_, okNW := currentLocations[nw]
	_, okN := currentLocations[n]
	_, okNE := currentLocations[ne]
	_, okE := currentLocations[e]
	_, okSE := currentLocations[se]
	_, okS := currentLocations[s]
	_, okSW := currentLocations[sw]
	_, okW := currentLocations[w]

	if !(okNW || okN || okNE || okE || okSE || okS || okSW || okW) {
		return true
	}

	switch d.Name {
	case "North":
		return okNW || okN || okNE
	case "South":
		return okSW || okS || okSE
	case "West":
		return okNW || okW || okSW
	case "East":
		return okNE || okE || okSE
	}
	return true
}

func (elf *Elf) simulateMove(d Direction, nextLocations map[string][]*Elf) {
	switch d.Name {
	case "North":
		loc := fmt.Sprintf("%d,%d", elf.Current.Row-1, elf.Current.Col)
		nextLocations[loc] = append(nextLocations[loc], elf)
	case "South":
		loc := fmt.Sprintf("%d,%d", elf.Current.Row+1, elf.Current.Col)
		nextLocations[loc] = append(nextLocations[loc], elf)
	case "West":
		loc := fmt.Sprintf("%d,%d", elf.Current.Row, elf.Current.Col-1)
		nextLocations[loc] = append(nextLocations[loc], elf)
	case "East":
		loc := fmt.Sprintf("%d,%d", elf.Current.Row, elf.Current.Col+1)
		nextLocations[loc] = append(nextLocations[loc], elf)
	}
}

func printGrid(minRow, maxRow, minCol, maxCol int, currentLocations map[string]string) {
	for i := minRow; i <= maxRow; i++ {
		for j := minCol; j <= maxCol; j++ {
			if _, ok := currentLocations[fmt.Sprintf("%d,%d", i, j)]; !ok {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}
