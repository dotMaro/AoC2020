package main

import (
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day7/input.txt")

	rules := parseRules(input)
	bag := rules.bagWithColor("shiny gold")
	count := bag.containedByCount()
	utils.Print("Task 1. There are %d bags that contain shiny gold bags.", count)
	containCount := bag.mustContain()
	utils.Print("Task 2. Shiny gold bags must contain %d other bags.", containCount)
}

func parseRules(input string) rules {
	lines := utils.SplitLine(input)
	rules := rules{
		bags: []*bag{},
	}
	for _, line := range lines {
		splitLine := strings.SplitN(line, " bags contain ", 2)
		bagColor := splitLine[0]
		alreadyAddedBag := rules.bagWithColor(bagColor)
		var curBag *bag
		if alreadyAddedBag == nil {
			curBag = &bag{
				color: bagColor,
			}
		} else {
			curBag = alreadyAddedBag
		}
		if !strings.HasSuffix(line, "no other bags.") {
			contents := strings.Split(splitLine[1], ", ")
			containedBags := make([]bagAmount, len(contents))
			for i, c := range contents {
				splitContent := strings.SplitN(c, " ", 2)
				count, err := strconv.Atoi(splitContent[0])
				if err != nil {
					panic(err)
				}
				color := strings.SplitN(splitContent[1], " bag", 2)[0]
				containedBag := rules.bagWithColor(color)
				if containedBag == nil {
					containedBag = &bag{
						color: color,
					}
					rules.bags = append(rules.bags, containedBag)
				}
				containedBag.appendContainedBy(curBag)

				containedBags[i] = bagAmount{
					bag:   containedBag,
					count: count,
				}
			}
			curBag.contains = containedBags

			// Only add new if it wasn't added already
			if alreadyAddedBag == nil {
				rules.bags = append(rules.bags, curBag)
			}
		}
	}

	return rules
}

type rules struct {
	bags []*bag
}

func (r rules) bagWithColor(c string) *bag {
	for _, b := range r.bags {
		if b.color == c {
			return b
		}
	}

	return nil
}

type bag struct {
	color       string
	contains    []bagAmount
	containedBy []*bag
}

func (b *bag) containsCount(c string) int {
	for _, bagAmount := range b.contains {
		if bagAmount.color == c {
			return bagAmount.count
		}
	}
	return 0
}

func (b *bag) appendContainedBy(o *bag) {
	if b.containedBy == nil || len(b.containedBy) == 0 {
		b.containedBy = []*bag{o}
	} else {
		b.containedBy = append(b.containedBy, o)
	}
}

func (b *bag) containedByCount() int {
	return b.containedByCountRecursive(make(map[string]struct{})) - 1 // -1 to not include the bag in question (b)
}

func (b *bag) containedByCountRecursive(counted map[string]struct{}) int {
	_, alreadyCounted := counted[b.color]
	if len(b.containedBy) == 0 {
		if alreadyCounted {
			return 0
		}
		counted[b.color] = struct{}{}
		return 1
	}

	count := 0
	for _, c := range b.containedBy {
		count += c.containedByCountRecursive(counted)
	}

	if !alreadyCounted {
		counted[b.color] = struct{}{}
		count++
	}
	return count
}

func (b *bag) mustContain() int {
	return b.mustContainRecursive() - 1 // -1 to not include the bag in question (b)
}

func (b *bag) mustContainRecursive() int {
	if len(b.contains) == 0 {
		return 1
	}
	totalCount := 1
	for _, c := range b.contains {
		totalCount += c.count * c.mustContainRecursive()
	}
	return totalCount
}

type bagAmount struct {
	*bag
	count int
}
