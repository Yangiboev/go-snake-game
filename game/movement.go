package game

// MoveUp moves the snake horizontally up
func (g *Game) MoveUp() {
	g.headX--
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.GetTail()
}

// MoveRight moves the snake vertically right
func (g *Game) MoveRight() {
	g.headY++
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.GetTail()
}

// MoveDown moves the snake horizontally down
func (g *Game) MoveDown() {
	g.headX++
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.GetTail()
}

// MoveLeft moves the snake vertically left
func (g *Game) MoveLeft() {
	g.headY--
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.GetTail()
}

// getTail fetches the latest location of the tail after a move
func (g *Game) GetTail() {
	if g.length == 1 {
		g.tailX = g.headX
		g.tailY = g.headY
		return
	}

	if g.tailX != 0 && g.board[g.tailX-1][g.tailY] == 1 {
		g.tailX--
		return
	}

	if g.tailX != g.boardHeight-1 && g.board[g.tailX+1][g.tailY] == 1 {
		g.tailX++
		return
	}

	if g.tailY != 0 && g.board[g.tailX][g.tailY-1] == 1 {
		g.tailY--
		return
	}

	if g.tailY != g.boardWidth-1 && g.board[g.tailX][g.tailY+1] == 1 {
		g.tailY++
		return
	}

}
