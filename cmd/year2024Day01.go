package cmd

import (
	"github.com/hnatekmarorg/adventOfCode2024/utils"
	"github.com/spf13/cobra"
	"log"
	"slices"
	"strconv"
)

func partOne(firstColumn []int, secondColumn []int) {
	log.Println("Part one calculating...")
	slices.Sort(firstColumn)
	slices.Sort(secondColumn)
	if len(firstColumn) != len(secondColumn) {
		panic("Must be same length")
	}
	totalDistance := 0
	for i := 0; i < len(firstColumn); i++ {
		distance := firstColumn[i] - secondColumn[i]
		if distance < 0 {
			distance *= -1
		}
		totalDistance += distance
	}
	log.Printf("%d\n", totalDistance)
}

func partTwo(firstColumn []int, secondColumn []int) {
	log.Println("Part two calculating...")
	//countCache := map[string]int{} // Optionally this could be optimized with cache
	result := 0
	for i := range firstColumn {
		result += firstColumn[i] * utils.CountOccurrences(firstColumn[i], secondColumn)
	}
	log.Printf("%d\n", result)
}

// year2024Day01Cmd represents the year2024Day01 command
var year2024Day01Cmd = &cobra.Command{
	Use:   "year2024Day01",
	Short: "First day of 2024",
	Long:  `First day of 2024 https://adventofcode.com/2024/day/1`,
	Run: func(cmd *cobra.Command, args []string) {
		data := utils.LoadSolution(args[0])
		firstColumn := utils.ConvertColumn(data, utils.StringConversionWrapper(strconv.Atoi), 0)
		secondColumn := utils.ConvertColumn(data, utils.StringConversionWrapper(strconv.Atoi), 1)
		partOne(firstColumn, secondColumn)
		firstColumn = utils.ConvertColumn(data, utils.StringConversionWrapper(strconv.Atoi), 0)
		secondColumn = utils.ConvertColumn(data, utils.StringConversionWrapper(strconv.Atoi), 1)
		partTwo(firstColumn, secondColumn)
	},
}

func init() {
	rootCmd.AddCommand(year2024Day01Cmd)

	// File path
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// year2024Day01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// year2024Day01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
