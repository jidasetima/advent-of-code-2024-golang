package day1

import (
	"errors"
	"testing"
)

// TestPart1 calls day1.Part1 with 2 lists of equal lengths with difference in each step being positive,
// checking for a valid return value
func TestPart1PositiveRanges(t *testing.T) {
	list1 := []int{2, 1, 3}
	list2 := []int{4, 2, 3}
	want := 3

	value, err := Part1(list1, list2)
	if err != nil {
		t.Errorf("Unexpected error occured: %d", err)
	} else if value != want {
		t.Errorf("Part1([]int {1,2,3}, []int {2,3,4}) = %d; want 3", value)
	}
}

// TestPart1 calls day1.Part1 with 2 lists of equal lengths with difference in each step being negative,
// checking for a valid return value
func TestPart1NegativeRanges(t *testing.T) {
	list1 := []int{7, 6, 5}
	list2 := []int{3, 2, 1}
	want := 12

	value, err := Part1(list1, list2)
	if err != nil {
		t.Errorf("Unexpected error occured: %d", err)
	} else if value != want {
		t.Errorf("Part1([]int {1,2,3}, []int {2,3,4}) = %d; want %d", value, want)
	}
}

func TestPart1ListUnequalLengthError(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{1, 2}
	want := ErrNonMatchingListLengths

	_, err := Part1(list1, list2)
	if !errors.Is(err, want) {
		t.Errorf("Error = %d, want %d", err, want)
	}
}

func TestPart2(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{1, 2, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5, 5}
	want := 47

	value := Part2(list1, list2)
	if value != want {
		t.Errorf("Part2([]int {1,2,3,4,5}, []int {1,2,2,2,3,4,4,4,5,5,5,5,5} = %d; want %d", value, want)
	}
}
