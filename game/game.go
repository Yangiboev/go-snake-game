package game

import (
	"math/rand"
	"time"
)

// Game stores all the game data
type Game struct {
	Direction string

	Round int
	State int

	board       [][]int8
	score       int
	length      int
	headX       int
	headY       int
	tailX       int
	tailY       int
	boardHeight int
	boardWidth  int
	foodX       int
	foodY       int
}

// Intialize the game
func NewGame() *Game {
	g := new(Game)

	g.boardHeight = 20
	g.boardWidth = 20
	g.Direction = "right"
	g.ResetGame()

	return g
}

//  resetGame resets to game to its initial stage
//  and assigns default values to game variables
func (g *Game) ResetGame() {
	g.board = make([][]int8, g.boardHeight)

	for i := 0; i < g.boardHeight; i++ {
		g.board[i] = make([]int8, g.boardWidth)

		for j := 0; j < g.boardWidth; j++ {
			g.board[i][j] = 0
		}
	}

	g.board[0][0] = 1
	g.Round = 0
	g.State = 1
	g.score = 0
	g.length = 1
	g.headX = 0
	g.headY = 0
	g.tailX = 0
	g.tailY = 0
	g.GetFood()
}

func (g *Game) GetFood() {
	rand.Seed(time.Now().UnixNano())

	i, j := g.GetRand()

	// 1 in game array means snake.
	if g.board[i][j] == 1 {
		// Check if food will spawn on snake
		g.GetFood()
	} else {
		g.board[i][j] = 2
		g.foodX = i
		g.foodY = j
	}
}

// CheckFood checks if snake head is on the food
func (g *Game) CheckFood() {
	if g.foodX == g.headX && g.foodY == g.headY {
		g.score++
		g.length++
		g.board[g.foodX][g.foodY] = 1

		switch g.Direction {

		case "up":
			g.tailX++
			g.board[g.tailX][g.tailY] = 1

		case "right":
			g.tailY--
			g.board[g.tailX][g.tailY] = 1

		case "down":
			g.tailX--
			g.board[g.tailX][g.tailY] = 1

		case "left":
			g.tailY++
			g.board[g.tailX][g.tailY] = 1

		default:
			panic("direction is wrong")
		}

		g.GetFood()
	}
}

// CheckMove checks if the move of snake is valid, it returns
// true if valid else false
func (g *Game) CheckMove() bool {
	switch g.Direction {

	case "up":
		if g.headX == 0 || g.board[g.headX-1][g.headY] == 1 {
			return false
		}
		return true

	case "right":
		if g.headY == g.boardHeight-1 || g.board[g.headX][g.headY+1] == 1 {
			return false
		}
		return true

	case "down":
		if g.headX == g.boardWidth-1 || g.board[g.headX+1][g.headY] == 1 {
			return false
		}
		return true

	case "left":
		if g.headY == 0 || g.board[g.headX][g.headY-1] == 1 {
			return false
		}
		return true

	default:
		panic("Direction is wrong")
	}

}

// min returns the minimum of board Width and Height
func (g *Game) Min() int {
	if g.boardHeight < g.boardWidth {
		return g.boardHeight
	}
	return g.boardWidth
}

// getRand returns two random integer between 0 and min
func (g *Game) GetRand() (int, int) {
	rand.Seed(time.Now().UnixNano())
	rand1 := rand.Intn(g.Min())

	rand.Seed(time.Now().UnixNano())
	rand2 := rand.Intn(g.Min())

	return rand2, rand1
}

// getBoardEnd returns coords of board edges
func (g *Game) GetBoardEnd() (int, int) {
	var boardEndX int = 2 + g.boardWidth*2
	var boardEndY int = 3

	return boardEndX + 2, boardEndY
}
