package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

type Directory struct {
	name    string
	size    int
	subdirs map[string]*Directory
	parent  *Directory
}

func createStructure(lines []string) *Directory {
	root := &Directory{"/", 0, nil, nil}
	cwd := root

	reChangeDirectory := regexp.MustCompile(`\$ cd (.*)$`)

	for _, line := range lines[1:] { // skipping the `cd /` as it's handled at root
		if line == "$ cd .." {
			cwd = cwd.parent
		} else if strings.HasPrefix(line, "$ cd") {
			directory := reChangeDirectory.FindStringSubmatch(line)
			if len(directory) != 2 {
				log.Fatalf("length mismsatch, %s => %v\n", line, directory)
			}

			nextDir, exists := cwd.subdirs[directory[1]]
			if !exists { // i'm thinking this is impossible given the input
				log.Fatalf("Failed to find %v\n", directory[1])
			}

			cwd = nextDir
		} else { // `ls` or the output of `ls`
			splitLine := strings.Split(line, " ")
			if !strings.HasPrefix(line, "dir") && !strings.HasPrefix(line, "$ ls") {
				size, err := strconv.Atoi(splitLine[0])
				if err != nil {
					log.Fatalf("error converting size: %v\n", err)
				}

				cwd.size += size
			} else if line != "$ ls" {
				if cwd.subdirs == nil {
					cwd.subdirs = make(map[string]*Directory)
				}

				cwd.subdirs[splitLine[1]] = &Directory{
					name:    splitLine[1],
					size:    0,
					subdirs: make(map[string]*Directory),
					parent:  cwd,
				}
			} else {
				continue
			}
		}
	}

	return root
}

func aggregateSizes(dir *Directory) int {
	totalSize := dir.size
	for _, subdir := range dir.subdirs {
		totalSize += aggregateSizes(subdir)
	}

	dir.size = totalSize

	// if dir.size <= 100_000 {
	// 	*sum += dir.size
	// }

	return totalSize
}

func p1ProcessLines(lines []string) (int, error) {
	root := createStructure(lines)
	aggregateSizes(root)

	stack := aoc.Stack[*Directory]{}
	stack.Push(root)

	sum := 0
	seen := make(map[*Directory]bool)
	for !stack.Empty() {
		cwd := stack.Pop()
		if seen[cwd] {
			continue
		}
		seen[cwd] = true

		if cwd.size < 100000 {
			sum += cwd.size
		}

		for _, dir := range cwd.subdirs {
			if !seen[dir] {
				stack.Push(dir)
			}
		}
	}

	// 1749646
	// 1340491 - to low
	// 1280098
	return sum, nil
}

func p2ProcessLines(lines []string) (int, error) {
	root := createStructure(lines)

	totalUsedSpace := aggregateSizes(root)
	totalUnusedSpace := 70_000_000 - totalUsedSpace
	requiredForUpdate := 30_000_000 - totalUnusedSpace
	if requiredForUpdate <= 0 {
		log.Fatalf("Required space already found...")
	}

	stack := aoc.Stack[*Directory]{}
	stack.Push(root)

	seen := make(map[*Directory]bool)

	smallestPossibleFolder := root
	for !stack.Empty() {
		cwd := stack.Pop()
		if seen[cwd] {
			continue
		}
		seen[cwd] = true

		if cwd.size >= requiredForUpdate && cwd.size <= smallestPossibleFolder.size {
			// fmt.Printf("setting smallest folder too: %v\n", smallestPossibleFolder)
			smallestPossibleFolder = cwd
		}

		for _, dir := range cwd.subdirs {
			if !seen[dir] {
				stack.Push(dir)
			}
		}
	}

	// 33688513 -- too high
	return smallestPossibleFolder.size, nil
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	sum, err := p1ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 1 Error: %e\n", err)
	}

	fmt.Printf("Part 1 Sum = %d\n", sum)

	sum, err = p2ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 2 Error: %e\n", err)
	}

	fmt.Printf("Part 2 Sum = %d\n", sum)
}
