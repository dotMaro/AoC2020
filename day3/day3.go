package main

import (
	"sync"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day3/input.txt")
	treeMap := newTreeMap(input)

	treesHit2 := treeMap.traverse(3, 1)
	utils.Print("Task 1. Trees hit %d", treesHit2)
	product := treeMap.multiTraverse([]slope{{1, 1}, {5, 1}, {7, 1}, {1, 2}}) * treesHit2
	// product := treeMap.concurrentMultiTraverse([]slope{{1, 1}, {5, 1}, {7, 1}, {1, 2}}) * treesHit2
	utils.Print("Task 2. Product %d", product)
}

func newTreeMap(input string) treeMap {
	splitInput := utils.SplitLine(input)

	treeMap := make([][]bool, len(splitInput))
	for y, line := range splitInput {
		treeMap[y] = make([]bool, len(splitInput[0]))
		for x, c := range line {
			treeMap[y][x] = c == '#'
		}
	}
	return treeMap
}

type treeMap [][]bool

// traverse on slope x, y and return the amount of trees hit.
func (m treeMap) traverse(x, y int) int {
	curPos := struct{ x, y int }{0, 0}
	treesHit := 0

	for curPos.y < len(m) {
		if m[curPos.y][curPos.x%len(m[0])] {
			treesHit++
		}
		curPos.x += x
		curPos.y += y
	}
	return treesHit
}

type slope struct{ x, y int }

// multitraverse traverses in multiple slopes and return the product of trees hit.
func (m treeMap) multiTraverse(slopes []slope) int {
	product := 1
	for _, s := range slopes {
		product *= m.traverse(s.x, s.y)
	}
	return product
}

func (m treeMap) concurrentTraverse(wg *sync.WaitGroup, c chan<- int, x, y int) {
	c <- m.traverse(x, y)
	wg.Done()
}

// multitraverse traverses in multiple slopes and return the product of trees hit.
func (m treeMap) concurrentMultiTraverse(slopes []slope) int {
	var (
		product    = 1
		slopeCount = len(slopes)
		wg         sync.WaitGroup
		c          = make(chan int, slopeCount)
	)
	wg.Add(slopeCount)
	for _, s := range slopes {
		go m.concurrentTraverse(&wg, c, s.x, s.y)
	}
	for i := 0; i < slopeCount; i++ {
		product *= <-c
	}
	wg.Wait()
	return product
}
