package main

import "testing"

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
