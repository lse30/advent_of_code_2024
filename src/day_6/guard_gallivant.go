package guard_gallivant

import (
	"advent_of_code_2024/src/utils"
	"fmt"
)

const guard = '^'

type coord struct {
	x int
	y int
}

func (c coord) Add(other coord) coord {
	return coord{
		x: c.x + other.x,
		y: c.y + other.y,
	}
}

type cachedMapState struct {
	x int
	y int
	d int // direction as an int
}

func turn(currentDirection coord) coord {
	if currentDirection.x == -1 {
		return coord{0, 1}
	}
	if currentDirection.y == 1 {
		return coord{1, 0}
	}
	if currentDirection.x == 1 {
		return coord{0, -1}
	} else {
		return coord{-1, 0}
	}
}

func canGoStraight(data []string, position coord) bool {
	if inBounds(data, position) {
		return data[position.x][position.y] != '#'
	}
	return false
}

func inBounds(data []string, position coord) bool {
	if position.x < 0 || position.x >= len(data) {
		return false
	}
	if position.y < 0 || position.y >= len(data[0]) {
		return false
	}
	return true
}

func updateMap(mapData []string, position coord, targetChar rune) []string {
	row := []rune(mapData[position.x])
	row[position.y] = targetChar
	mapData[position.x] = string(row)
	return mapData
}

func moveGuard(data []string, guardPosition coord, direction coord) []string {
	nextPosition := guardPosition.Add(direction)
	if inBounds(data, nextPosition) {
		if canGoStraight(data, nextPosition) {
			data := updateMap(data, nextPosition, 'X')
			moveGuard(data, nextPosition, direction)
		} else {
			newDirection := turn(direction)
			moveGuard(data, guardPosition, newDirection)
		}
	}
	return data
}

func directionCoordToInt(d coord) int {
	if d.x == -1 {
		return 1
	}
	if d.y == 1 {
		return 2
	}
	if d.x == 1 {
		return 3
	} else {
		return 4
	}
}

func inCache(cache []cachedMapState, newCache cachedMapState) bool {
	for i := 0; i < len(cache); i++ {
		if cache[i].x == newCache.x && cache[i].y == newCache.y && cache[i].d == newCache.d {
			return true
		}
	}
	return false
}

func endsInLoop(data []string, guardPosition coord, direction coord, caches []cachedMapState) bool {
	nextCache := cachedMapState{guardPosition.x, guardPosition.y, directionCoordToInt(direction)}
	if inCache(caches, nextCache) {
		return true
	} else {
		caches = append(caches, nextCache)
	}

	nextPosition := guardPosition.Add(direction)

	if inBounds(data, nextPosition) {
		if canGoStraight(data, nextPosition) {
			return endsInLoop(data, nextPosition, direction, caches)
		} else {
			newDirection := turn(direction)
			return endsInLoop(data, guardPosition, newDirection, caches)
		}
	}
	return false
}

func countNewObstructions(data []string) int {
	output := 0
	guardStart := findGuardPosition(data)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == '.' {
				data = updateMap(data, coord{i, j}, '#')
				if endsInLoop(data, guardStart, coord{-1, 0}, []cachedMapState{}) {
					//printMap(data)
					output++
				}
				data = updateMap(data, coord{i, j}, '.')
			}
		}
	}
	return output

}

func findGuardPosition(data []string) coord {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == guard {
				return coord{i, j}
			}
		}
	}
	return coord{1, 1}
}

func countX(mapData []string) int {
	output := 0
	for i := 0; i < len(mapData); i++ {
		fmt.Println(mapData[i])
		for j := 0; j < len(mapData[i]); j++ {
			if mapData[i][j] == 'X' || mapData[i][j] == '^' {
				output++
			}
		}
	}
	return output
}

func SolvePartOne(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	guardPosition := findGuardPosition(data)
	finalMap := moveGuard(data, guardPosition, coord{-1, 0})
	return countX(finalMap)
}

func SolvePartTwo(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	count := countNewObstructions(data)
	return count
}
