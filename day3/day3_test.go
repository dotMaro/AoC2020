package main

import (
	"testing"

	"github.com/dotMaro/AoC2020/utils"
)

const input = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestTraverse(t *testing.T) {
	m := newTreeMap(input)
	treesHit := m.traverse(3, 1)
	if treesHit != 7 {
		t.Errorf("Should hit 7 trees, but hit %d", treesHit)
	}
}

const slopeCount = 6

func BenchmarkMultiTraverse(b *testing.B) {
	input := utils.ReadFile("input.txt")
	treeMap := newTreeMap(input)
	slopes := make([]slope, slopeCount)
	for i := 0; i < slopeCount; i++ {
		slopes[i] = slope{1, 1}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		treeMap.multiTraverse(slopes)
	}
}

func BenchmarkConcurrentMultiTraverse(b *testing.B) {
	input := utils.ReadFile("input.txt")
	treeMap := newTreeMap(input)
	slopes := make([]slope, slopeCount)
	for i := 0; i < slopeCount; i++ {
		slopes[i] = slope{1, 1}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		treeMap.concurrentMultiTraverse(slopes)
	}
}
