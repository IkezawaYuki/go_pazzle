package main

import (
	"fmt"
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


func checkedReset(){
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


func main() {
	var cursorX = 0
	var cursorY = 0
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++{
			num := rand.Intn(7)
			cell[x][y] = block[num]
		}
	}

	var selectedX = -1
	var selectedY = -1
	locked := false



	for{
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()


		// 落下処理
		for y := HEIGHT-2; y >= 0; y-- {
			for x := 0; x < WIDTH; x++ {
		    	if cell[y][x] != "-" && cell[y+1][x] == "-"{
					 cell[y+1][x] = cell[y][x]
					 cell[y][x] = "-"
				}
			}
		}

		// 画面描画処理
		for y := 0; y < HEIGHT; y++ {
			for x := 0; x < WIDTH; x++{
				if cursorX == x && cursorY == y{
					fmt.Print("[" +cell[y][x]+"]")
				} else{
					fmt.Print(" " + cell[y][x]+" ")
				}
			}
			if selectedY == y {
				fmt.Print("←")
			}
			fmt.Println()
		}

		for x := 0; x < WIDTH; x++{
			if selectedX == x{
				fmt.Print("↑ ")
			}else {
				fmt.Print("  ")
			}
		}
		fmt.Println()


		var word string
		fmt.Scan(&word)

		if !locked {
			// キーボード入力処理
			switch word {
			case "s":
				cursorY += 1
				if cursorY > 7{
					cursorY = 7
				}
			case "w":
				cursorY -= 1
				if cursorY < 0{
					cursorY = 0
				}
			case "a":
				cursorX -= 1
				if cursorX < 0{
					cursorX = 0
				}
			case "d":
				cursorX += 1
				if cursorX > 7{
					cursorX = 7
				}
			case "z":
				if selectedX < 0{
					selectedX = cursorX
					selectedY = cursorY
				}else{
					cell[selectedY][selectedX], cell[cursorY][cursorX] = cell[cursorY][cursorX], cell[selectedY][selectedX]

					// 連結確認処理
					checkedReset()
					for y := 0; y < HEIGHT; y++ {
						for x := 0; x < WIDTH; x++{
							n := getConnectedBlockCount(x, y, cell[y][x], 0)
							if n >= 3{
								eraseConnectedBlock(x, y, cell[y][x])
							}
						}
					}
					selectedX, selectedY = -1, -1
					locked = true
				}
			}
		}
	}
}

