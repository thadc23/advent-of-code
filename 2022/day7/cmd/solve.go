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

	"github.com/spf13/cobra"
)

type File struct {
	Name string
	Size int
}

type Folder struct {
	Name    string
	Size    int
	Files   []*File
	Folders []*Folder
	Parent  *Folder
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

	root := &Folder{Name: "/"}
	var curr *Folder

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		switch line[0] {
		case "$":
			switch line[1] {
			case "cd":
				if line[2] == "/" {
					curr = root
				} else if line[2] == ".." {
					curr = curr.Parent
				} else {
					newFolder := &Folder{Name: line[2], Parent: curr}
					curr.Folders = append(curr.Folders, newFolder)
					curr = newFolder
				}
			}
		case "dir":
		default:
			size, _ := strconv.Atoi(line[0])
			curr.Files = append(curr.Files, &File{Size: size, Name: line[1]})
			curr.Size = curr.Size + size
			parent := curr.Parent
			for parent != nil {
				parent.Size = parent.Size + size
				parent = parent.Parent
			}
		}
	}

	if part2 {
		spaceNeeded := 30000000 - (70000000 - root.Size)
		allFolders := findAllFolders(root)

		sort.Slice(allFolders, func(i, j int) bool {
			return allFolders[i].Size < allFolders[j].Size
		})

		for _, f := range allFolders {
			if f.Size > spaceNeeded {
				return f.Size
			}
		}
	}
	return calc(root)

}

func calc(folder *Folder) int {
	size := 0
	for _, f := range folder.Folders {
		size += calc(f)
	}
	if folder.Size < 100000 {
		return size + folder.Size
	}
	return size
}

func findAllFolders(folder *Folder) []*Folder {
	var folders []*Folder
	for _, f := range folder.Folders {
		folders = append(folders, findAllFolders(f)...)
	}
	return append(folders, folder.Folders...)
}
