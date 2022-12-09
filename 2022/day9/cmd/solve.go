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

	"github.com/spf13/cobra"
)

type Point struct {
	X int
	Y int
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
	points := make(map[string]string)

	var rope []*Point
	if part2 {
		rope = make([]*Point, 10)
	} else {
		rope = make([]*Point, 2)
	}
	for i := 0; i < len(rope); i++ {
		rope[i] = &Point{X: 0, Y: 0}
	}

	head := rope[len(rope)-1]
	recordSpot(rope[0], points)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		moves, _ := strconv.Atoi(line[1])
		for i := 0; i < moves; i++ {
			switch line[0] {
			case "U":
				head.Y = head.Y + 1
			case "D":
				head.Y = head.Y - 1
			case "R":
				head.X = head.X + 1
			case "L":
				head.X = head.X - 1
			}

			for j := len(rope) - 1; j > 0 && !closeEnough(rope[j], rope[j-1]); j-- {
				front := rope[j]
				back := rope[j-1]

				if front.X > back.X {
					back.X += 1
				}

				if front.Y > back.Y {
					back.Y += 1
				}

				if front.Y < back.Y {
					back.Y -= 1
				}

				if front.X < back.X {
					back.X -= 1
				}

			}
			recordSpot(rope[0], points)
		}
	}
	return len(points)
}

func closeEnough(curr, next *Point) bool {
	return math.Abs(float64(next.X)-float64(curr.X)) <= 1 && math.Abs(float64(next.Y)-float64(curr.Y)) <= 1
}

func recordSpot(point *Point, points map[string]string) {
	x := strconv.Itoa(point.X)
	y := strconv.Itoa(point.Y)

	points[x+y] = "X"
}
