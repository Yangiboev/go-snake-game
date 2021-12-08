package main

import (
	"github.com/Yangiboev/go-snake-game/game"
	"github.com/nsf/termbox-go"
)

/**
* * Author: Dilmurod Yangiboev
* ? Email: dilmurod.yangiboev@gmail.com
* ? Linkedin: www.linkedin.com/in/dilmurod-yangiboev
 */

func main() {
	game := game.NewGame()
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	eventQueue := make(chan termbox.Event)

	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	game.Render()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowLeft:
					game.Direction = "left"
					if game.CheckMove() && game.State == 1 {
						game.Round++
						game.MoveLeft()
						game.CheckFood()
					} else {
						game.State = 0
					}
				case ev.Key == termbox.KeyArrowRight:
					game.Direction = "right"
					if game.CheckMove() && game.State == 1 {
						game.Round++
						game.MoveRight()
						game.CheckFood()
					} else {
						game.State = 0
					}
				case ev.Key == termbox.KeyArrowUp:
					game.Direction = "up"
					if game.CheckMove() && game.State == 1 {
						game.Round++
						game.MoveUp()
						game.CheckFood()
					} else {
						game.State = 0
					}
				case ev.Key == termbox.KeyArrowDown:
					game.Direction = "down"

					if game.CheckMove() && game.State == 1 {
						game.Round++
						game.MoveDown()
						game.CheckFood()
					} else {
						game.State = 0
					}

				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
		default:
			game.Render()
		}
	}
}
