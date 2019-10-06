package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

func NewPlay() *Player{
	return &Player{
		error: "",
		locked: false,
		cursorX: 0,
		cursorY: 0,
		selectedX: -1,
		selectedY: -1,
	}
}

type Player struct {
	error string
	locked bool
	cursorX int
	cursorY int
	selectedX int
	selectedY int
}

func (p *Player) Display(board *Board){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	if p.error != ""{
		fmt.Println(p.error)
	}


	p.locked = false

	// 落下処理
	for y := HEIGHT-2; y >= 0; y-- {
		for x := 0; x < WIDTH; x++ {
			if board.cell[y][x] != "-" && board.cell[y+1][x] == "-"{
				board.cell[y+1][x] = board.cell[y][x]
				board.cell[y][x] = "-"
				p.locked = true
			}
		}
	}

	// 補充処理
	for x := 0; x < WIDTH; x++{
		if board.cell[0][x] == "-"{
			board.cell[0][x] = block[rand.Intn(7)]
			p.locked = true
		}
	}

	// 画面描画処理
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++{
			if p.cursorX == x && p.cursorY == y{
				fmt.Print("[" +board.cell[y][x]+"]")
			} else{
				fmt.Print(" " + board.cell[y][x]+" ")
			}
		}
		if p.selectedY == y {
			fmt.Print("←")
		}
		fmt.Println()
	}

	for x := 0; x < WIDTH; x++{
		if p.selectedX == x{
			fmt.Print(" ↑ ")
		}else {
			fmt.Print("  ")
		}
	}
	fmt.Println()

	if !p.locked{
		board.EraseConnectedBlockAll(p)
	}
}
