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
	"math"
	"os"
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

		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		if part2 {
			solvePart2(scanner)
		} else {
			solvePart1(scanner)
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

func solvePart1(scanner *bufio.Scanner) {
	count := 0
	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), " | ")
		nums := strings.Split(slice[1], " ")
		for _, num := range nums {
			switch len(num) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	fmt.Println(count)
}

func solvePart2(scanner *bufio.Scanner) {
	var finalTally float64
	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), " | ")
		signals := strings.Split(slice[0], " ")
		nums := strings.Split(slice[1], " ")
		numMap := make(map[int]string)
		for i := 0; i < 10; i++ {
			numMap[i] = ""
		}
		numMap, signals = findEasyOnes(signals, numMap)
		numMap, signals = findSix(signals, numMap)
		numMap, signals = findNineAndZero(signals, numMap)
		numMap, signals = findThree(signals, numMap)
		numMap, signals = findTwoAndFive(signals, numMap)

		for i, num := range nums {
			val := sort(num)
			for k, v := range numMap {
				if val == v {
					finalTally += float64(k) * math.Pow(10, float64(3-i))
				}
			}
		}
	}
	fmt.Println(int(finalTally))
}

func findEasyOnes(signals []string, numMap map[int]string) (map[int]string, []string) {
	for i := 0; i < len(signals); i++ {
		sorted := sort(signals[i])
		switch len(sorted) {
		case 2:
			numMap[1] = sorted
			signals = append(signals[:i], signals[i+1:]...)
			i--
		case 3:
			numMap[7] = sorted
			signals = append(signals[:i], signals[i+1:]...)
			i--
		case 4:
			numMap[4] = sorted
			signals = append(signals[:i], signals[i+1:]...)
			i--
		case 7:
			numMap[8] = sorted
			signals = append(signals[:i], signals[i+1:]...)
			i--
		}
	}
	return numMap, signals
}

func findSix(signals []string, numMap map[int]string) (map[int]string, []string) {
	for i := 0; i < len(signals); i++ {
		sorted := sort(signals[i])
		switch len(sorted) {
		case 6:
			if !containsAll(sorted, numMap[1]) {
				numMap[6] = sorted
				signals = append(signals[:i], signals[i+1:]...)
				return numMap, signals
			}
		}
	}
	return numMap, signals
}

func findNineAndZero(signals []string, numMap map[int]string) (map[int]string, []string) {
	for i := 0; i < len(signals); i++ {
		sorted := sort(signals[i])
		switch len(sorted) {
		case 6:
			if containsAll(sorted, numMap[4]) {
				numMap[9] = sorted
				signals = append(signals[:i], signals[i+1:]...)
				i--
			} else {
				numMap[0] = sorted
				signals = append(signals[:i], signals[i+1:]...)
				i--
			}
		}
	}
	return numMap, signals
}

func sort(num string) string {
	str := []rune(num)
	for x := range str {
		y := x + 1
		for y = range str {
			if str[x] < str[y] {
				str[x], str[y] = str[y], str[x]
			}
		}
	}
	return string(str)
}

func findThree(signals []string, numMap map[int]string) (map[int]string, []string) {
	for i := 0; i < len(signals); i++ {
		sorted := sort(signals[i])
		switch len(sorted) {
		case 5:
			if containsAll(sorted, numMap[1]) {
				numMap[3] = sorted
				signals = append(signals[:i], signals[i+1:]...)
				return numMap, signals
			}
		}
	}
	return numMap, signals
}

func findTwoAndFive(signals []string, numMap map[int]string) (map[int]string, []string) {

	signal1 := signals[0]
	signal2 := signals[1]

	var finalSignal1, finalSignal2 string

	for _, x := range signal1 {
		if !strings.Contains(signal2, string(x)) {
			finalSignal1 += string(x)
		}
	}
	for _, x := range signal2 {
		if !strings.Contains(signal1, string(x)) {
			finalSignal2 += string(x)
		}
	}

	if containsAll(numMap[6], finalSignal1) {
		numMap[5] = sort(signal1)
		numMap[2] = sort(signal2)
		return numMap, signals
	} else {
		numMap[2] = sort(signal1)
		numMap[5] = sort(signal2)
		return numMap, signals
	}

}

func containsAll(str, substr string) bool {
	for _, x := range substr {
		if !strings.Contains(str, string(x)) {
			return false
		}
	}
	return true
}
