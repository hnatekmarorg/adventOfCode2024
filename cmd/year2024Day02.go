package cmd

import (
	"fmt"
	"github.com/hnatekmarorg/adventOfCode2024/utils"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func checkAscending(row []int) bool {
	for i := 1; i < len(row); i++ {
		if row[i-1] > row[i] {
			return false
		}
	}
	return true
}

func checkDescending(row []int) bool {
	for i := 1; i < len(row); i++ {
		if row[i-1] < row[i] {
			return false
		}
	}
	return true
}

func checkDiff(row []int, maxDiff int, minDiff int) bool {
	for i := 1; i < len(row); i++ {
		diff := row[i-1] - row[i]
		if diff < 0 {
			diff *= -1
		}
		if diff > maxDiff || diff < minDiff {
			return false
		}
	}
	return true
}

func year2024Day02PartOne(data [][]string) {
	log.Println("Part one calculating...")
	safe := 0
	for row := range data {
		convertedRow := utils.ConvertRow(data, utils.StringConversionWrapper(strconv.Atoi), row)

		if (checkAscending(convertedRow) || checkDescending(convertedRow)) && checkDiff(convertedRow, 3, 1) {
			safe++
		}
	}
	log.Printf("%d\n", safe)
}

func year2024Day02PartTwo(data [][]string) {
	log.Println("Part two calculating...")
	safe := 0
	for row := range data {
		convertedRow := utils.ConvertRow(data, utils.StringConversionWrapper(strconv.Atoi), row)
		if (checkAscending(convertedRow) || checkDescending(convertedRow)) && checkDiff(convertedRow, 3, 1) {
			fmt.Println("OK", convertedRow)
			safe++
			continue
		}
		for i := 0; i < len(convertedRow); i++ {
			copyRow := make([]int, len(convertedRow))
			copy(copyRow, convertedRow)
			// Create a new slice without the ith element
			convertedWithoutElement := append(copyRow[:i], copyRow[i+1:]...)

			fmt.Println("Converted[", i, "]", convertedRow, "->", convertedWithoutElement)
			if (checkAscending(convertedWithoutElement) || checkDescending(convertedWithoutElement)) && checkDiff(convertedWithoutElement, 3, 1) {
				safe++
				break
			}
		}
		fmt.Println("NOK", convertedRow)
	}
	log.Printf("%d\n", safe)
}

// year2024Day01Cmd represents the year2024Day01 command
var year2024Day02Cmd = &cobra.Command{
	Use:   "year2024Day02",
	Short: "Second day of 2024",
	Long:  `Second day of 2024 https://adventofcode.com/2024/day/2`,
	Run: func(cmd *cobra.Command, args []string) {
		data := utils.LoadSolution(args[0])
		year2024Day02PartOne(data)
		year2024Day02PartTwo(data)
	},
}

func init() {
	rootCmd.AddCommand(year2024Day02Cmd)

	// File path
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// year2024Day01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// year2024Day01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
