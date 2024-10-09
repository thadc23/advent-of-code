/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

type Mapping struct {
	destination int
	source      int
	step_range  int
}

func calc(scanner *bufio.Scanner, part2 bool) int {

	seeds := findSeeds(scanner)
	scanner.Scan()
	scanner.Scan()
	seed_to_soil := createMapping(scanner)
	scanner.Scan()
	soil_to_fertilizer := createMapping(scanner)
	scanner.Scan()
	fertilizer_to_water := createMapping(scanner)
	scanner.Scan()
	water_to_light := createMapping(scanner)
	scanner.Scan()
	light_to_temp := createMapping(scanner)
	scanner.Scan()
	temp_to_humidity := createMapping(scanner)
	scanner.Scan()
	humidity_to_location := createMapping(scanner)

	min_location := math.MaxInt64

	if part2 {
		seedMappings := expandSeeds(seeds)
		for i := 0; i < math.MaxInt64; i++ {

			humidity := findSpotReverse(i, humidity_to_location)
			temperature := findSpotReverse(humidity, temp_to_humidity)
			light := findSpotReverse(temperature, light_to_temp)
			water := findSpotReverse(light, water_to_light)
			fertilizer := findSpotReverse(water, fertilizer_to_water)
			soil := findSpotReverse(fertilizer, soil_to_fertilizer)
			seed := findSpotReverse(soil, seed_to_soil)

			if mappingContains(seed, seedMappings) {
				return i
			}
		}
	}

	for i := 0; i < math.MaxInt64; i++ {
		humidity := findSpotReverse(i, humidity_to_location)
		temperature := findSpotReverse(humidity, temp_to_humidity)
		light := findSpotReverse(temperature, light_to_temp)
		water := findSpotReverse(light, water_to_light)
		fertilizer := findSpotReverse(water, fertilizer_to_water)
		soil := findSpotReverse(fertilizer, soil_to_fertilizer)
		seed := findSpotReverse(soil, seed_to_soil)

		if contains(seed, seeds) {
			return i
		}
	}

	return min_location

}

func mappingContains(seed int, mappings []Mapping) bool {
	for _, m := range mappings {
		if seed >= m.source && seed < m.source+m.step_range {
			return true
		}
	}
	return false
}

func contains(val int, vals []int) bool {
	for _, v := range vals {
		if v == val {
			return true
		}
	}
	return false
}

func expandSeeds(seeds []int) []Mapping {
	var new_seeds []Mapping

	for i := 0; i < len(seeds); i = i + 2 {
		new_seeds = append(new_seeds, Mapping{source: seeds[i], step_range: seeds[i+1]})
	}
	return new_seeds
}

func findSpotReverse(spot int, mappings []Mapping) int {
	for _, m := range mappings {
		if spot >= m.destination && spot < m.destination+m.step_range {
			return m.source + (spot - m.destination)
		}
	}
	return spot
}

func createMapping(scanner *bufio.Scanner) []Mapping {
	var mappings []Mapping

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fields := strings.Fields(line)
		values := convertStringsToInts(fields)
		mappings = append(mappings, Mapping{destination: values[0], source: values[1], step_range: values[2]})
	}
	return mappings
}

func findSeeds(scanner *bufio.Scanner) []int {
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ": ")
	return convertStringsToInts(strings.Fields(parts[1]))
}

func convertStringsToInts(strArray []string) []int {
	var intArray []int

	for _, str := range strArray {
		// Parse string to int
		num, _ := strconv.Atoi(str)
		// Append the parsed int to the new array
		intArray = append(intArray, num)
	}

	return intArray
}
