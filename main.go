package main

import "fmt"

func main() {
	var images [][]byte
	for i := 0; i < 9; i++ {
		images = append(images, []byte{})
	}
	puzzle := GeneratePuzzle(images, 3, false)
	fmt.Println(puzzle.IsSolvable())
	c := PuzzleController{
		puzzleGame:   PuzzleGame{},
		puzzleStatus: 0,
		puzzle:       puzzle,
		PuzzleGameID: "1",
		stepsTaken:   0,
	}

	c.tapTile(puzzle.tiles[2])
	fmt.Println(c.puzzle)
}