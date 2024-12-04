package day1

import (
	"slices"
)

func Part1(list1 []int, list2 []int) int {
	sum := 0

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		if difference := list2[i] - list1[i]; difference < 0 {
			sum = sum + (difference * -1)
		} else {
			sum = sum + difference
		}
	}

	return sum
}

func Part2(list1 []int, list2 []int) int {
	list2FrequencyMap := make(map[int]int)
	similarityScore := 0

	for _, value := range list2 {
		list2FrequencyMap[value] = list2FrequencyMap[value] + 1
	}

	for _, value := range list1 {
		similarityScore += value * list2FrequencyMap[value]
	}

	return similarityScore
}
