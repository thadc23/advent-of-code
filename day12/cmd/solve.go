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
	"os"
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

		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		input := make(map[string][]string)

		for scanner.Scan() {
			line := scanner.Text()
			slice := strings.Split(line, "-")

			if _, ok := input[slice[0]]; ok {
				input[slice[0]] = append(input[slice[0]], slice[1])
			} else {
				input[slice[0]] = []string{slice[1]}
			}

			if _, ok := input[slice[1]]; ok {
				input[slice[1]] = append(input[slice[1]], slice[0])
			} else {
				input[slice[1]] = []string{slice[0]}
			}

		}

		if part2 {
			solvePart2(input)
		} else {
			solvePart1(input)
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

func solvePart1(input map[string][]string) {
	_, paths := findPaths("start", input, "", map[string]string{})
	fmt.Println(len(paths))
}
func solvePart2(input map[string][]string) {
	_, paths := findPaths2("start", input, "", map[string]string{})
	fmt.Println(len(paths))
}

func findPaths(node string, tree map[string][]string, path string, paths map[string]string) (string, map[string]string) {
	// fmt.Println(node)
	startPath := path
	if node == "end" {
		finalPath := path
		finalPath += node
		paths[finalPath] = finalPath
		return path, paths
	} else {
		startPath += node + ","
	}

	for _, connection := range tree[node] {
		if !((IsLower(connection) && occurrences(path, connection) == 1) || connection == "start") {
			// fmt.Printf("%s -> %s, continue path=%s\n", node, connection, path)
			path, paths = findPaths(connection, tree, startPath, paths)
		}
	}
	return startPath, paths
}

func findPaths2(node string, tree map[string][]string, path string, paths map[string]string) (string, map[string]string) {
	// fmt.Println(node)
	startPath := path
	if node == "end" {
		finalPath := path
		finalPath += node
		paths[finalPath] = finalPath
		return path, paths
	} else {
		startPath += node + ","
	}

	for _, connection := range tree[node] {
		if !(((IsLower(connection) && occurrences(path, connection) == 2) || (IsLower(connection) && occurrences(path, connection) == 1 && hasVisitedTooMany(startPath))) || connection == "start") {
			fmt.Printf("%s -> %s, continue path=%s\n", node, connection, path)
			path, paths = findPaths2(connection, tree, startPath, paths)
		}
	}
	return startPath, paths
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func occurrences(path string, connection string) int {
	count := 0
	slice := strings.Split(path, ",")
	for _, x := range slice {
		if x == connection {
			count++
		}
	}
	return count
}

func hasVisitedTooMany(path string) bool {

	caves := make(map[string]int)

	parts := strings.Split(path, ",")

	for _, x := range parts {
		if IsLower(x) {
			caves[x]++
		}
	}

	for _, x := range caves {
		if x > 1 {
			return true
		}
	}

	return false

}
