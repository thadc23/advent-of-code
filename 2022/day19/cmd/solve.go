/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

type RobotRequirement struct {
	Ore int
}

type OreRobot struct {
	RobotRequirement
}
type ClayRobot struct {
	RobotRequirement
}
type ObsidianRobot struct {
	RobotRequirement
	Clay int
}
type GeodeRobot struct {
	RobotRequirement
	Obsidian int
}

type Blueprint struct {
	Index int

	OreRobot
	ClayRobot
	ObsidianRobot
	GeodeRobot

	Ore      int
	Clay     int
	Obsidian int
	Geode    int

	OreRobots      []int
	ClayRobots     []int
	ObsidianRobots []int
	GeodeRobots    []int
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
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fmt.Println(score(scanner, part2))
}

func score(scanner *bufio.Scanner, part2 bool) int {

	r := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

	var blueprints []*Blueprint
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindStringSubmatch(line)

		blueprints = append(blueprints, newBlueprint(match[1:]))
	}

	total := 0
	for _, b := range blueprints {
		for i := 0; i < 24; i++ {
			newOre := len(b.OreRobots)
			newClay := len(b.ClayRobots)
			newObsidian := len(b.ObsidianRobots)
			newGeode := len(b.GeodeRobots)
			if canBuildGeode(b) {
				fmt.Println("built Geode Robot")
				b.Ore -= b.GeodeRobot.Ore
				b.Obsidian -= b.GeodeRobot.Obsidian
				b.GeodeRobots = append(b.GeodeRobots, 1)
			}
			if canBuildObsidian(b) {
				fmt.Println("built Obsidian Robot")
				b.Ore -= b.ObsidianRobot.Ore
				b.Clay -= b.ObsidianRobot.Clay
				b.ObsidianRobots = append(b.ObsidianRobots, 1)
			}
			if canBuildClay(b) {
				fmt.Println("built Clay Robot")
				b.Ore -= b.ClayRobot.Ore
				b.ClayRobots = append(b.ClayRobots, 1)
			}
			if canBuildOre(b) {
				fmt.Println("built Ore Robot")
				b.Ore -= b.OreRobot.Ore
				b.OreRobots = append(b.OreRobots, 1)
			}

			b.Ore += newOre
			fmt.Printf("%d ore robot(s) collected %d ore. new value = %d\n", len(b.OreRobots), newOre, b.Ore)
			b.Clay += newClay
			b.Obsidian += newObsidian
			b.Geode += newGeode

		}
		total += (b.Geode * b.Index)
		fmt.Println(b.Geode)
	}
	return total
}

func newBlueprint(vals []string) *Blueprint {
	index, _ := strconv.Atoi(vals[0])
	oreOre, _ := strconv.Atoi(vals[1])
	clayOre, _ := strconv.Atoi(vals[2])
	obsidianOre, _ := strconv.Atoi(vals[3])
	obsidianClay, _ := strconv.Atoi(vals[4])
	geodeOre, _ := strconv.Atoi(vals[5])
	geodeObsidian, _ := strconv.Atoi(vals[6])

	return &Blueprint{
		Index:         index,
		OreRobot:      OreRobot{RobotRequirement{Ore: oreOre}},
		ClayRobot:     ClayRobot{RobotRequirement{Ore: clayOre}},
		ObsidianRobot: ObsidianRobot{RobotRequirement: RobotRequirement{Ore: obsidianOre}, Clay: obsidianClay},
		GeodeRobot:    GeodeRobot{RobotRequirement: RobotRequirement{Ore: geodeOre}, Obsidian: geodeObsidian},
		OreRobots:     []int{1},
	}
}

func canBuildGeode(blueprint *Blueprint) bool {
	return blueprint.Ore >= blueprint.GeodeRobot.Ore && blueprint.Obsidian >= blueprint.GeodeRobot.Obsidian
}
func canBuildObsidian(blueprint *Blueprint) bool {
	return blueprint.Ore >= blueprint.ObsidianRobot.Ore && blueprint.Clay >= blueprint.ObsidianRobot.Clay
}
func canBuildClay(blueprint *Blueprint) bool {
	return blueprint.Ore >= blueprint.ClayRobot.Ore
}
func canBuildOre(blueprint *Blueprint) bool {
	return blueprint.Ore >= blueprint.OreRobot.Ore
}
