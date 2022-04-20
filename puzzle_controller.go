package gopuzzlegame

import "math/rand"

type PuzzleStatus int

const (
	PuzzleStatusIncomplete = iota
	PuzzleStatusComplete
	PuzzleStatusReachedStepLimit
)

type PuzzleController struct {
	//PuzzleGame   PuzzleGame
	PuzzleStatus PuzzleStatus
	Puzzle       *Puzzle
	StepsTaken   int32
	Steps		 int32
}

func (p *PuzzleController) TapTile(tile *Tile) {
	if p.PuzzleStatus == PuzzleStatusIncomplete && p.Puzzle.IsTileMovable(tile) {
		mutablePuzzle := Puzzle{Tiles: p.Puzzle.Tiles}
		var tiles []*Tile
		p.Puzzle = mutablePuzzle.MoveTiles(tile, tiles)
		p.Puzzle.Sort()
		p.StepsTaken++
		if p.Puzzle.IsComplete() {
			p.PuzzleStatus = PuzzleStatusComplete
		} else if p.StepsTaken == p.Steps {
			p.PuzzleStatus = PuzzleStatusReachedStepLimit
		}
	}
}

func GeneratePuzzle(images []Image, size int, shuffle bool) *Puzzle {
	var correctPositions []Position
	var currentPositions []Position
	whitespacePosition := Position{
		X: size,
		Y: size,
	}

	for y := 1; y <= size; y++ {
		for x := 1; x <= size; x++ {
			if x == size && y == size {
				correctPositions = append(correctPositions, whitespacePosition)
				currentPositions = append(currentPositions, whitespacePosition)
			} else {
				position := Position{
					X: x,
					Y: y,
				}
				correctPositions = append(correctPositions, position)
				currentPositions = append(currentPositions, position)
			}
		}
	}

	if shuffle {
		rand.Shuffle(len(currentPositions), func(i, j int) {
			currentPositions[i], currentPositions[j] = currentPositions[j], currentPositions[i]
		})
	}

	tiles := getTileListFromPositions(size, images, correctPositions, currentPositions)
	puzzle := &Puzzle{Tiles: tiles}

	if shuffle {
		for !puzzle.IsSolvable() || puzzle.GetNumberOfCorrectTiles() != 0 {
			rand.Shuffle(len(currentPositions), func(i, j int) {
				currentPositions[i], currentPositions[j] = currentPositions[j], currentPositions[i]
			})
			puzzle = &Puzzle{Tiles: getTileListFromPositions(size, images, correctPositions, currentPositions)}
		}
	}

	return puzzle
}

func getTileListFromPositions(size int, images []Image, correctPositions, currentPositions []Position) []*Tile {
	whitespacePosition := Position{
		X: size,
		Y: size,
	}

	var result []*Tile
	for i := 1; i <= size*size; i++ {
		if i == size*size {
			result = append(result, &Tile{
				Image:           images[i-1],
				Value:           i,
				CorrectPosition: whitespacePosition,
				CurrentPosition: currentPositions[i-1],
				IsWhitespace:    true,
			})
		} else {
			result = append(result, &Tile{
				Image:           images[i-1],
				Value:           i,
				CorrectPosition: correctPositions[i-1],
				CurrentPosition: currentPositions[i-1],
			})
		}
	}

	return result
}
