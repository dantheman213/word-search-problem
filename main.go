package main

// Post: https://reddit.com/r/mildlyinfuriating/comments/s7778u/been_looking_for_meow_for_30_minutes_while_in_a/
// Word search image: https://i.redd.it/x1mx6zytbic81.jpg

import (
	"fmt"
	"strings"
)

func main() {
	jumble := [20][15]string {
		{"F", "C", "X", "B", "A", "L", "L", "O", "O", "N", "W", "U", "B", "R", "D"},
		{"I", "K", "H", "A", "Y", "B", "T", "R", "X", "M", "O", "O", "E", "S", "Z"},
		{"E", "M", "M", "I", "S", "N", "O", "D", "F", "T", "I", "R", "I", "J", "R"},
		{"L", "U", "E", "A", "C", "T", "O", "U", "H", "V", "N", "O", "H", "T", "A"},
		{"D", "I", "C", "R", "C", "K", "X", "C", "J", "Y", "K", "O", "V", "R", "B"},
		{"O", "C", "X", "A", "W", "O", "R", "K", "G", "H", "K", "S", "S", "A", "B"},
		{"H", "O", "R", "S", "E", "E", "W", "L", "U", "K", "W", "T", "K", "I", "I"},
		{"F", "T", "Z", "W", "M", "D", "V", "I", "E", "J", "I", "E", "E", "L", "T"},
		{"C", "H", "I", "R", "P", "D", "B", "N", "X", "B", "G", "R", "X", "E", "X"},
		{"B", "A", "A", "A", "G", "F", "Y", "G", "S", "K", "O", "Z", "T", "R", "S"},
		{"S", "F", "T", "J", "I", "L", "E", "N", "F", "L", "N", "S", "N", "Y", "Y"},
		{"U", "G", "B", "Q", "M", "R", "R", "N", "S", "H", "E", "E", "P", "W", "N"},
		{"Q", "J", "A", "N", "E", "A", "P", "L", "C", "V", "M", "T", "R", "E", "E"},
		{"U", "V", "A", "V", "B", "K", "L", "L", "R", "E", "O", "K", "K", "Y", "U"},
		{"A", "V", "X", "J", "M", "I", "F", "A", "A", "O", "U", "D", "H", "M", "S"},
		{"C", "P", "G", "O", "M", "B", "H", "Z", "M", "N", "S", "I", "O", "G", "H"},
		{"K", "I", "D", "D", "O", "M", "Y", "L", "L", "C", "E", "I", "G", "G", "G"},
		{"L", "G", "N", "Q", "F", "E", "I", "A", "U", "Q", "L", "E", "I", "T", "F"},
		{"U", "I", "W", "C", "P", "A", "S", "T", "U", "R", "E", "E", "Y", "X", "Q"},
		{"W", "C", "Z", "C", "P", "K", "P", "X", "H", "Q", "N", "K", "U", "A", "K"},
	}

	wordlist := []string {
		"AIRPLANE",
		"BAA",
		"BALLOON",
		"BARN",
		"CAT",
		"CHICK",
		"CHIRP",
		"COW",
		"DOG",
		"DUCKLING",
		"EGGS",
		"FARMER",
		"FENCE",
		"FIELD",
		"HARVEST",
		"HAY",
		"HORSE",
		"MEOW",
		"MOO",
		"MOUSE",
		"NEIGH",
		"OINK",
		"PASTURE",
		"PIG",
		"QUACK",
		"RABBIT",
		"ROOSTER",
		"SHEEP",
		"TRACTOR",
		"TRAILER",
		"TREE",
		"WINDMILL",
	}

	for _, word := range wordlist {
		letters := strings.Split(word, "")
		checkHorizontal(letters, jumble)
		checkVertical(letters, jumble)
		checkDiagonalForwardDown(letters, jumble)
		checkDiagonalForwardUp(letters, jumble)
	}
}

func checkHorizontal(word []string, grid [20][15]string) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]) - len(word) + 1; x++ {
			if grid[y][x] == word[0] {
				searchPattern := grid[y][x:x+len(word)]
				if isStrArrayEqual(word, searchPattern) {
					fmt.Printf("found horizontal word %s at x: %d y: %d\n", word, x, y)
				}
			}
		}
	}
}

func checkVertical(word []string, grid [20][15]string) {
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid) - len(word) + 1; y++ {
			if grid[y][x] == word[0] {
				searchPattern := []string{}
				for y1 := y; y1 < y + len(word); y1++ {
					searchPattern = append(searchPattern, grid[y1][x])
				}

				if isStrArrayEqual(word, searchPattern) {
					fmt.Printf("found vertical word %s at x: %d y: %d\n", word, x, y)
				}
			}
		}
	}
}

func checkDiagonalForwardDown(word []string, grid [20][15]string) {
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid) - len(word) + 1; y++ {
			if grid[y][x] == word[0] {
				x1 := x
				y1 := y
				searchPattern := []string{}
				for x1 < len(grid[y]) && y1 < len(grid) {
					searchPattern = append(searchPattern, grid[y1][x1])
					x1 += 1
					y1 += 1

					if len(searchPattern) == len(word) {
						break
					}
				}

				if isStrArrayEqual(word, searchPattern) {
					fmt.Printf("found diagonal word %s at x: %d y: %d\n", word, x, y)
				}
			}
		}
	}
}

func checkDiagonalForwardUp(word []string, grid [20][15]string) {
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid) - len(word) + 1; y++ {
			if grid[y][x] == word[0] {
				x1 := x
				y1 := y
				searchPattern := []string{}
				for x1 < len(grid[y]) && y1 < len(grid) {
					searchPattern = append(searchPattern, grid[y1][x1])
					x1 += 1
					y1 += 1

					if len(searchPattern) == len(word) {
						break
					}
				}

				if isStrArrayEqual(word, searchPattern) {
					fmt.Printf("found diagonal word %s at x: %d y: %d\n", word, x, y)
				}
			}
		}
	}
}

func isStrArrayEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

