package main

import (
	"fmt"
	"strings"
)

func makeIsland(s string) [][]byte {
	rows := strings.Split(s, "\n")
	island := make([][]byte, len(rows))
	numCols := len(rows[0])
	for i := 0; i < numCols; i++ {
		island[i] = make([]byte, len(rows[0]))
	}
	for i := range rows {
		copy(island[i], []byte(rows[i]))
	}
	return island
}

type posn struct {
	x int
	y int
}

func numIslands1(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var mark func(int, int)
	yMax := len(grid) - 1
	xMax := len(grid[0]) - 1
	mapp := make(map[posn]struct{})
	mark = func(x, y int) {
		if _, ok := mapp[posn{x, y}]; ok || grid[y][x] != '1' {
			return
		}
		mapp[posn{x, y}] = struct{}{}
		if x+1 <= xMax {
			mark(x+1, y)
		}
		if y+1 <= yMax {
			mark(x, y+1)
		}
		if x-1 >= 0 {
			mark(x-1, y)
		}
		if y-1 >= 0 {
			mark(x, y-1)
		}
	}
	islands := 0
	for y := range grid {
		for x := range grid[y] {
			if _, ok := mapp[posn{x, y}]; !ok && grid[y][x] == '1' {
				islands++
			}
			mark(x, y)
		}
	}
	return islands
}

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var mark func(int, int)
	yMax := len(grid) - 1
	xMax := len(grid[0]) - 1
	mark = func(x, y int) {
		if grid[y][x] != '1' {
			return
		}
		grid[y][x] = 0
		if x+1 <= xMax {
			mark(x+1, y)
		}
		if y+1 <= yMax {
			mark(x, y+1)
		}
		if x-1 >= 0 {
			mark(x-1, y)
		}
		if y-1 >= 0 {
			mark(x, y-1)
		}
	}
	islands := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '1' {
				islands++
			}
			mark(x, y)
		}
	}
	return islands
}

func main() {
	island := "" +
		"1001\n" +
		"1100\n" +
		"1011\n" +
		"1001"
		/* 	island := "" +
		"11\n" +
		"01" */
	islandBytes := makeIsland(island)
	fmt.Println(numIslands2(islandBytes))
}
