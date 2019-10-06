package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
)

const (
	HEIGHT = 8
	WIDTH = 8
	BLOCK_TYPE_MAX = 7
	CELL_TYPE_NONE = "・"
	CELL_TYPE_BLOCK_0 = "○"
	CELL_TYPE_BLOCK_1 = "△"
	CELL_TYPE_BLOCK_2 = "□"
	CELL_TYPE_BLOCK_3 = "●"
	CELL_TYPE_BLOCK_4 = "▲"
	CELL_TYPE_BLOCK_5 = "■"
	CELL_TYPE_BLOCK_6 = "☆"
)

var block = []string{
	"○",
	"△",
	"□",
	"●",
	"▲",
	"■",
	"☆",
}

var cell = [][]string{
	{"#","#","#","#","#","#","#","#"},
	{"#","#","#","#","#","#","#","#"},
	{"#","#","#","#","#","#","#","#"},
	{"#","#","#","#","#","#","#","#"},
	{"#","#","#","#","#","#","#","#"},
	{"#","#","#","#","#","#","#","#"},
	{"#","#","#","#","#","#","#","#"},
	{"#","#","#","#","#","#","#","#"},
}

var checked = [][]bool{
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
}


func ResetChecked(){
	checked = [][]bool{
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false},
	}
}



func getConnectedBlockCount(x int, y int, cellType string, count int) int{
	if x < 0 || x >= WIDTH || y < 0 || y >= HEIGHT || checked[y][x] || cell[y][x] == "-" || cell[y][x] != cellType{
		return count
	}

	count += 1
	checked[y][x] = true

	count = getConnectedBlockCount(x, y+1, cellType, count)
	count = getConnectedBlockCount(x, y-1, cellType, count)
	count = getConnectedBlockCount(x+1, y, cellType, count)
	count = getConnectedBlockCount(x-1, y, cellType, count)

	return count
}

func eraseConnectedBlock(x int, y int, cellType string){
	if x < 0 || x >= WIDTH || y < 0 || y >= HEIGHT || cell[y][x] == "-" || cell[y][x] != cellType{
		return
	}

	cell[y][x] = "-"
	eraseConnectedBlock(x-1, y, cellType)
	eraseConnectedBlock(x, y-1, cellType)
	eraseConnectedBlock(x+1, y, cellType)
	eraseConnectedBlock(x, y+1, cellType)
}

func eraseConnectedBlockAll(player *Player){
	ResetChecked()
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++{
			n := getConnectedBlockCount(x, y, cell[y][x], 0)
			if n >= 3{
				eraseConnectedBlock(x, y, cell[y][x])
				player.locked = true
			}
		}
	}
}

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

func (p *Player) display(){
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
			if cell[y][x] != "-" && cell[y+1][x] == "-"{
				cell[y+1][x] = cell[y][x]
				cell[y][x] = "-"
				p.locked = true
			}
		}
	}

	// 補充処理
	for x := 0; x < WIDTH; x++{
		if cell[0][x] == "-"{
			cell[0][x] = block[rand.Intn(7)]
			p.locked = true
		}
	}

	// 画面描画処理
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++{
			if p.cursorX == x && p.cursorY == y{
				fmt.Print("[" +cell[y][x]+"]")
			} else{
				fmt.Print(" " + cell[y][x]+" ")
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
		eraseConnectedBlockAll(p)
	}
}


func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++{
			cell[x][y] = block[rand.Intn(7)]
		}
	}

	play := NewPlay()


	for{
		play.display()

		if !play.locked {
			var word string
			fmt.Scan(&word)
			play.error = ""

			// キーボード入力処理
			switch word {
			case "s":
				play.cursorY += 1
				if play.cursorY > 7{
					play.cursorY = 7
				}
			case "w":
				play.cursorY -= 1
				if play.cursorY < 0{
					play.cursorY = 0
				}
			case "a":
				play.cursorX -= 1
				if play.cursorX < 0{
					play.cursorX = 0
				}
			case "d":
				play.cursorX += 1
				if play.cursorX > 7{
					play.cursorX = 7
				}
			case "z":
				if play.selectedX < 0 {
					play.selectedX = play.cursorX
					play.selectedY = play.cursorY
				}else{

					// 隣接したブロックを洗濯しているかどうか
					distance := math.Abs(float64(play.selectedX) - float64(play.cursorX)) +
						math.Abs(float64(play.selectedY) - float64(play.cursorY))
					if distance == 1.0{
						// 入れ替え処理
						cell[play.selectedY][play.selectedX], cell[play.cursorY][play.cursorX] = cell[play.cursorY][play.cursorX], cell[play.selectedY][play.selectedX]
						// 連結確認処理

						eraseConnectedBlockAll(play)

						play.selectedX, play.selectedY = -1, -1
						play.locked = true
					}else{
						play.error = "隣接していないブロック同士は入れ替えることが出来ません。"
					}
				}
			}
		}
	}
}

