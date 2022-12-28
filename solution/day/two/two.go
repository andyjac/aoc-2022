package two

import (
	"strconv"
	"strings"
)

type Two struct{}

const (
	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"
	playerRock       = "X"
	playerPaper      = "Y"
	playerScissors   = "Z"
	playerNeedsLoss  = "X"
	playerNeedsDraw  = "Y"
	playerNeedsWin   = "Z"
)

const (
	drawPts = 3
	winPts  = 6
)

var shapePointMap = map[string]int{
	opponentRock:     1,
	opponentPaper:    2,
	opponentScissors: 3,
	playerRock:       1,
	playerPaper:      2,
	playerScissors:   3,
}

var shapeNeedMap = map[string]map[string]string{
	opponentRock: {
		playerNeedsLoss: opponentScissors,
		playerNeedsDraw: opponentRock,
		playerNeedsWin:  opponentPaper,
	},
	opponentPaper: {
		playerNeedsLoss: opponentRock,
		playerNeedsDraw: opponentPaper,
		playerNeedsWin:  opponentScissors,
	},
	opponentScissors: {
		playerNeedsLoss: opponentPaper,
		playerNeedsDraw: opponentScissors,
		playerNeedsWin:  opponentRock,
	},
}

var shapeWinMap = map[string]map[string]bool{
	opponentRock: {
		playerRock:     false,
		playerPaper:    false,
		playerScissors: true,
	},
	opponentPaper: {
		playerRock:     true,
		playerPaper:    false,
		playerScissors: false,
	},
	opponentScissors: {
		playerRock:     false,
		playerPaper:    true,
		playerScissors: false,
	},
	playerRock: {
		opponentRock:     false,
		opponentPaper:    false,
		opponentScissors: true,
	},
	playerPaper: {
		opponentRock:     true,
		opponentPaper:    false,
		opponentScissors: false,
	},
	playerScissors: {
		opponentRock:     false,
		opponentPaper:    true,
		opponentScissors: false,
	},
}

type CalcGameResult func(opponentMove string, playerMove string) int

func calcGameResultPart1(opponentShape string, playerShape string) int {
	var matchResultPts int

	opponentWins := shapeWinMap[opponentShape][playerShape]
	playerWins := shapeWinMap[playerShape][opponentShape]
	shapePts := shapePointMap[playerShape]

	if !opponentWins && !playerWins {
		matchResultPts = drawPts
	} else if playerWins {
		matchResultPts = winPts
	}

	return shapePts + matchResultPts
}

func calcGameResultPart2(opponentShape string, playerNeed string) int {
	var matchResultPoints int

	shapeToPlay := shapeNeedMap[opponentShape][playerNeed]
	shapePts := shapePointMap[shapeToPlay]

	if playerNeed == playerNeedsDraw {
		matchResultPoints = drawPts
	}
	if playerNeed == playerNeedsWin {
		matchResultPoints = winPts
	}

	return shapePts + matchResultPoints
}

func getGameResult(input string, calcGameResult CalcGameResult) (string, error) {
	lines := strings.Split(input, "\n")
	var total int

	for _, l := range lines {
		if l != "" {
			shapes := strings.Split(l, " ")
			pts := calcGameResult(shapes[0], shapes[1])
			total += pts
		}
	}

	return strconv.Itoa(total), nil
}

func PartOne(input string) (string, error) {
	return getGameResult(input, calcGameResultPart1)
}

func PartTwo(input string) (string, error) {
	return getGameResult(input, calcGameResultPart2)
}
