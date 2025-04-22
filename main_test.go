package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkWhole(b *testing.B) {
	var partSizes = []int{2, 5, 10, 20, 50}
	var counts = []int{2, 5, 10, 20, 50, 200, 1000, 4000}

	for _, partSize := range partSizes {
		b.Run("partSize="+strconv.Itoa(partSize), func(b *testing.B) {
			for _, count := range counts {
				b.Run("count="+strconv.Itoa(count), func(b *testing.B) {
					input := createRandomInput(partSize, count)
					b.Run("method=Merge", func(b *testing.B) {
						for b.Loop() {
							_ = Merge(input)
						}
					})
					b.Run("method=Basic", func(b *testing.B) {
						for b.Loop() {
							_ = Basic(input)
						}
					})
					b.Run("method=Simple", func(b *testing.B) {
						for b.Loop() {
							Merge(input)
						}
					})
					b.Run("method=Build", func(b *testing.B) {
						for b.Loop() {
							_ = Build(input)
						}
					})
				})
			}
		})

	}
}

func createRandomInput(maxPartSize, count int) []string {
	random := rand.New(rand.NewSource(17))

	ret := make([]string, count)
	for i := range count {
		var str []rune
		size := random.Intn(maxPartSize)
		for range size {
			str = append(str, rune('a'+random.Intn(int('z'-'a'))))
		}
		ret[i] = string(str)
	}
	return ret
}
