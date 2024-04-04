package SPADE

type Vanilla struct {
}

// NewVanilla returns new instantiation of vanilla search methods with "nil" values
func NewVanilla() *Vanilla {
	return &Vanilla{}
}

// queryTotalNum counting the total number of occurrences for a value
func (vanilla *Vanilla) queryTotalNum(data []int, value int) int {
	count := 0
	for _, item := range data {
		if item == value {
			count++
		}
	}
	return count
}

// queryNumRep count the number of repeat for a value in a sequence
// shows the duration of the query value and the number of Transition to that value
func (vanilla *Vanilla) queryNumRep(data []int, value int) map[int]int {
	digitCounts := make(map[int]int)
	j := 0
	for i, item := range data {
		if item == value {
			digitCounts[j]++
		} else {
			j = i + 1
		}
	}

	return digitCounts
}
