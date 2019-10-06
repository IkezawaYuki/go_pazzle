package main

import (
	"fmt"
	"math"
	"math/rand"
)


func main() {

	play := NewPlay()
	board := NewBorad()

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++{
			board.cell[x][y] = block[rand.Intn(BLOCK_TYPE_MAX)]
		}
	}

	for{
		play.Display(board)

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
					if distance == 0{
						// 同じ箇所を選択した場合はリセット
						play.selectedX = -1
						play.selectedY = -1
					} else if distance == 1.0 {
						// 入れ替え処理
						board.cell[play.selectedY][play.selectedX], board.cell[play.cursorY][play.cursorX] = board.cell[play.cursorY][play.cursorX], board.cell[play.selectedY][play.selectedX]
						// 連結確認処理

						board.EraseConnectedBlockAll(play)

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

