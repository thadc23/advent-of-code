/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var row map[float64]string = make(map[float64]string)

	for scanner.Scan() {
		line := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return r == rune('=') || unicode.IsLetter(r) || unicode.IsSpace(r) || r == rune(',') || r == rune(':')
		})

		sensorX, _ := strconv.ParseFloat(line[0], 64)
		sensorY, _ := strconv.ParseFloat(line[1], 64)
		beaconX, _ := strconv.ParseFloat(line[2], 64)
		beaconY, _ := strconv.ParseFloat(line[3], 64)

		if beaconY == 2000000 {
			row[beaconX] = "B"
		}

		distance := manhattanDistance(sensorX, beaconX, sensorY, beaconY)

		if math.Abs(float64(beaconY)-float64(sensorY)) < distance {

			startX := sensorX - (distance - math.Abs(float64(2000000)-float64(sensorY)))
			endX := sensorX + (distance - math.Abs(float64(2000000)-float64(sensorY)))

			for i := startX; i <= endX; i++ {
				if _, ok := row[i]; !ok {
					row[i] = "#"
				}
			}

		}

	}

	for k, v := range row {
		if v == "B" {
			delete(row, k)
		}
	}

	return len(row)
}

func manhattanDistance(x1, x2, y1, y2 float64) float64 {
	return math.Abs(x1-x2) + math.Abs(y1-y2)
}
