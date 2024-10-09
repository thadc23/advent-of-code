/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

type Hand struct {
	cards string
	bid   int
	rank  int
}

var powers map[string]int = make(map[string]int)

func solve(part2 bool) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fmt.Println(calc(scanner, part2))
}

func calc(scanner *bufio.Scanner, part2 bool) int {
	powers["2"] = 2
	powers["3"] = 3
	powers["4"] = 4
	powers["5"] = 5
	powers["6"] = 6
	powers["7"] = 7
	powers["8"] = 8
	powers["9"] = 9
	powers["T"] = 10
	powers["J"] = 11
	powers["Q"] = 12
	powers["K"] = 13
	powers["A"] = 14

	if part2 {
		powers["J"] = 1
	}

	total := 0
	var hands []Hand
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		bid, _ := strconv.Atoi(line[1])
		hands = append(hands, Hand{cards: line[0], bid: bid})
	}

	var fiveOfAKind []Hand
	var fourOfAKind []Hand
	var fullHouse []Hand
	var threeOfAKind []Hand
	var twoPair []Hand
	var onePair []Hand
	var highCard []Hand

	for _, h := range hands {
		var cardCounts map[string]int = make(map[string]int)
		for _, c := range strings.Split(h.cards, "") {
			val, ok := cardCounts[c]
			if ok {
				cardCounts[c] = val + 1
			} else {
				cardCounts[c] = 1
			}
		}
		keys := sortByValue(cardCounts)

		switch cardCounts[keys[0]] {
		case 5:
			fiveOfAKind = append(fiveOfAKind, h)
		case 4:
			if part2 && (keys[0] == "J" || keys[1] == "J") {
				fiveOfAKind = append(fiveOfAKind, h)
			} else {
				fourOfAKind = append(fourOfAKind, h)
			}
		case 3:
			if cardCounts[keys[1]] == 2 {
				if part2 && (keys[0] == "J" || keys[1] == "J") {
					fiveOfAKind = append(fiveOfAKind, h)
				} else {
					fullHouse = append(fullHouse, h)
				}
			} else {
				if part2 && (keys[0] == "J" || keys[1] == "J" || keys[2] == "J") {
					fourOfAKind = append(fourOfAKind, h)
				} else {
					threeOfAKind = append(threeOfAKind, h)
				}
			}
		case 2:
			if cardCounts[keys[1]] == 2 {
				if part2 && len(keys) == 3 && keys[2] == "J" {
					fullHouse = append(fullHouse, h)
				} else if part2 && (keys[0] == "J" || keys[1] == "J") {
					fourOfAKind = append(fourOfAKind, h)
				} else {
					twoPair = append(twoPair, h)
				}
			} else {
				if part2 && (keys[0] == "J" || keys[1] == "J" || keys[2] == "J" || keys[3] == "J") {
					threeOfAKind = append(threeOfAKind, h)
				} else {
					onePair = append(onePair, h)
				}
			}
		case 1:
			if part2 && (keys[0] == "J" || keys[1] == "J" || keys[2] == "J" || keys[3] == "J" || keys[4] == "J") {
				onePair = append(onePair, h)
			} else {
				highCard = append(highCard, h)
			}
		}
	}

	fiveOfAKind = sortHands(fiveOfAKind)
	fourOfAKind = sortHands(fourOfAKind)
	fullHouse = sortHands(fullHouse)
	threeOfAKind = sortHands(threeOfAKind)
	twoPair = sortHands(twoPair)
	onePair = sortHands(onePair)
	highCard = sortHands(highCard)

	rank := rankHands(1, highCard)
	rank = rankHands(rank, onePair)
	rank = rankHands(rank, twoPair)
	rank = rankHands(rank, threeOfAKind)
	rank = rankHands(rank, fullHouse)
	rank = rankHands(rank, fourOfAKind)
	_ = rankHands(rank, fiveOfAKind)

	total += scoreHands(fiveOfAKind)
	total += scoreHands(fourOfAKind)
	total += scoreHands(fullHouse)
	total += scoreHands(threeOfAKind)
	total += scoreHands(twoPair)
	total += scoreHands(onePair)
	total += scoreHands(highCard)

	// fmt.Println(fiveOfAKind)
	// fmt.Println(fourOfAKind)
	// fmt.Println(fullHouse)
	// fmt.Println(threeOfAKind)
	// fmt.Println(twoPair)
	// fmt.Println(onePair)
	// fmt.Println(highCard)

	return total
}

func scoreHands(hands []Hand) int {
	score := 0
	for _, h := range hands {
		score += (h.bid * h.rank)
	}
	return score
}

func sortByValue(m map[string]int) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys
}

func rankHands(rank int, hands []Hand) int {
	for i := 0; i < len(hands); i++ {
		hands[i].rank = rank
		rank++
	}
	return rank
}

func sortHands(hands []Hand) []Hand {
	sort.SliceStable(hands, func(i, j int) bool {
		left := strings.Split(hands[i].cards, "")
		right := strings.Split(hands[j].cards, "")

		for i, l := range left {
			r := right[i]
			if powers[l] == powers[r] {
				continue
			}
			return powers[r] > powers[l]
		}
		return false
	})
	return hands
}
