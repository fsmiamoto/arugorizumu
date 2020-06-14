package direct_addressing

// All keys must be withing [0,MaxKey)
const MaxKey = 10000

type Report struct {
	NumOfKeys            int
	NumOfDistinctKeys    int
	MostFrequentKey      int
	MostFrequentKeyCount int
}

// This function uses direct addressing to solve the counting problem
func NewReport(keys []int, maxKey int) *Report {
	distinctKeys := 0
	mostFrequentKey := -1

	counts := make([]int, maxKey)

	for i := range keys {
		key := keys[i]

		if counts[key] == 0 {
			distinctKeys++
		}

		if mostFrequentKey == -1 {
			mostFrequentKey = key
		}

		counts[key]++

		if counts[key] >= counts[mostFrequentKey] {
			mostFrequentKey = key
		}
	}

	return &Report{
		NumOfKeys:            len(keys),
		NumOfDistinctKeys:    distinctKeys,
		MostFrequentKey:      mostFrequentKey,
		MostFrequentKeyCount: counts[mostFrequentKey],
	}
}
