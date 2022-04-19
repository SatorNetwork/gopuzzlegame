package main

type PuzzleGame struct {
	id string
	episodeID string
	prizePool float64
	rewards float64
	bonusRewards float64
	xSize int
	steps int
	stepsTaken int
	status int
	result int
	image string
}

type PuzzleGameStatus int

const (
	PuzzleGameStatusNotStarted = iota
	PuzzleGameStatusInProgress
	PuzzleGameStatusFinished
)

type PuzzleGameResult int

const (
	PuzzleGameResultNotFinished = iota
	PuzzleGameResultUserWon
	PuzzleGameResultUserLost
)
