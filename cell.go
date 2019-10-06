package main

var block = []string{
	"○",
	"△",
	"□",
	"●",
	"▲",
	"■",
	"☆",
}

const (
	HEIGHT = 8
	WIDTH = 8
	BLOCK_TYPE_MAX = 7
)


type Board struct {
	cell [][]string
	checked [][] bool
}

func (b *Board) ResetChecked(){
	for i := 0; i < len(b.checked); i++{
		for j := 0; j < len(b.checked[i]); j++{
			b.checked[i][j] = false
		}
	}
}

func NewBorad() *Board{
	return &Board{
		cell: [][]string{
			{"#","#","#","#","#","#","#","#"},
			{"#","#","#","#","#","#","#","#"},
			{"#","#","#","#","#","#","#","#"},
			{"#","#","#","#","#","#","#","#"},
			{"#","#","#","#","#","#","#","#"},
			{"#","#","#","#","#","#","#","#"},
			{"#","#","#","#","#","#","#","#"},
			{"#","#","#","#","#","#","#","#"},
		},
		checked: [][]bool{
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
		},
	}
}


func (b *Board) GetConnectedBlockCount(x int, y int, cellType string, count int) int{
	if x < 0 || x >= WIDTH || y < 0 || y >= HEIGHT || b.checked[y][x] || b.cell[y][x] == "-" || b.cell[y][x] != cellType{
		return count
	}

	count += 1
	b.checked[y][x] = true

	count = b.GetConnectedBlockCount(x, y+1, cellType, count)
	count = b.GetConnectedBlockCount(x, y-1, cellType, count)
	count = b.GetConnectedBlockCount(x+1, y, cellType, count)
	count = b.GetConnectedBlockCount(x-1, y, cellType, count)

	return count
}

func (b *Board) EraseConnectedBlock(x int, y int, cellType string){
	if x < 0 || x >= WIDTH || y < 0 || y >= HEIGHT || b.cell[y][x] == "-" || b.cell[y][x] != cellType{
		return
	}

	b.cell[y][x] = "-"
	b.EraseConnectedBlock(x-1, y, cellType)
	b.EraseConnectedBlock(x, y-1, cellType)
	b.EraseConnectedBlock(x+1, y, cellType)
	b.EraseConnectedBlock(x, y+1, cellType)
}

func (b *Board) EraseConnectedBlockAll(player *Player){
	b.ResetChecked()
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++{
			n := b.GetConnectedBlockCount(x, y, b.cell[y][x], 0)
			if n >= 3{
				b.EraseConnectedBlock(x, y, b.cell[y][x])
				player.locked = true
			}
		}
	}
}
