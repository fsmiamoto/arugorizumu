package direct_addressing

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func BenchmarkDirectAddressing1K(b *testing.B) {
	keys := readKeysFromFile(b, "randInt1K.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewReport(keys, MaxKey)
	}
}

func BenchmarkDirectAddressing10K(b *testing.B) {
	keys := readKeysFromFile(b, "randInt10K.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewReport(keys, MaxKey)
	}
}

func BenchmarkDirectAddressing100K(b *testing.B) {
	keys := readKeysFromFile(b, "randInt100K.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewReport(keys, MaxKey)
	}
}

func BenchmarkDirectAddressing1M(b *testing.B) {
	keys := readKeysFromFile(b, "randInt1M.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewReport(keys, MaxKey)
	}
}

func readKeysFromFile(b *testing.B, filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		b.Fatal(err)
	}

	var keys []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, err := strconv.Atoi(scanner.Text())
		if err != nil {
			b.Fatal(err)
		}
		keys = append(keys, key)
	}

	return keys
}
