package SPADE

import "strconv"

type Vanilla struct {
}

// NewVanilla returns new instantiation of vanilla search method with "nil" values
func NewVanilla() *Vanilla {
	return &Vanilla{}
}

// countOccurrences uses Hash Table (for Unsorted Arrays) to count the number of all occurrences,
// This method works well for unsorted arrays and has a time complexity of O(d),
// where d is the number of digits in the number
func (vanilla *Vanilla) countOccurrences(data []int) map[int]int {
	digitCounts := make(map[int]int)

	for _, num := range data {
		numStr := strconv.Itoa(num)
		for _, digitRune := range numStr {
			digit, _ := strconv.Atoi(string(digitRune))
			digitCounts[digit]++
		}
	}

	return digitCounts
}

// countDirect uses loop for counting the number of occurrences for specific value,
// has complexity of O(n), where n is the len of data
func (vanilla *Vanilla) countDirect(data []int, value int) int {
	count := 0
	for i := range data {
		if i == value {
			count++
		}
	}
	return count
}
