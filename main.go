package main

import (
	"fmt"
	"strings"
)

/*
row:2 |
row:1 |
row:0 |---------
		col:0 col:1 col:2
*/

//数组第一维代表1行row
type World [COL][ROW]*Grid

//地图 8*8个格子，格子状态：0:无障碍物，1:有障碍物
func CreateWorld() (world *World) {
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			world[i][j] = &Grid{
				X: j,
				Y: i,
				H: 0,
				G: 0,
				S: StateFlat,
			}
		}
	}
	//开始点
	world[0][0].S = StateStart
	//目标点
	world[ROW-1][COL/2].S = StateTarget
	return
}

const world1 string =
`
....T...
........
........
B.BBB...
........
.......B
........
S.......
`

func LoadWorld(world string) (*World, bool) {
	w := strings.TrimSpace(world)

	ret := &World{}
	tmp := strings.Split(w, "\n")
	for i, row := range tmp {
		for j, v := range row {
			ret[COL-1-i][j] = &Grid{
				X: COL-1-i,
				Y: j,
				S: string(v),
			}
		}
	}
	return ret, true
}

func main() {
	fmt.Println("Hello, A star!")
	//world := CreateWorld()
	world, ok := LoadWorld(world1)
	if !ok {
		return
	}
	PrintWorld(world)
}
