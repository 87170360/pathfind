package main

import (
	"fmt"
	"math"
	"strings"
)

//格子
type Grid struct {
	S string //状态
	X int    //坐标
	Y int    //坐标
	H int    //路径增量H值, 使用曼哈顿距离方法
}

func (g *Grid) isStart() bool {
	return g.S == StateStart
}

func (g *Grid) isTarget() bool {
	return g.S == StateTarget
}

//数组第一维代表1行row
type World struct {
	grids  [COL][ROW]*Grid
	start  *Grid
	target *Grid
}

func (this *World) LoadWorld(world string) bool {
	w := strings.TrimSpace(world)

	tmp := strings.Split(w, "\n")
	for i, row := range tmp {
		for j, v := range row {
			c := COL - 1 - i
			grid := &Grid{
				S: string(v),
				Y: c,
				X: j,
			}

			this.grids[c][j] = grid
			if grid.isStart() {
				this.start = grid
			}
			if grid.isTarget() {
				this.target = grid
			}
		}
	}
	return true
}

func (this *World) getGridByPox(x, y int) *Grid {
	if y >= ROW || x >= COL || x < 0 || y < 0 {
		return nil
	}
	return this.grids[y][x]
}

func (this *World) Neighbors(posX, posY int) []*Grid {
	ret := make([]*Grid, 0, 8)
	offset := [8][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {1, 1}, {-1, -1}, {1, -1}}
	//上，下，左，右，上左，上右，下左，下右
	for _, v := range offset {
		x := posX + v[0]
		y := posY + v[1]

		n := this.getGridByPox(x, y)
		if n == nil {
			continue
		}
		ret = append(ret, n)
	}

	return ret
}

func (this *World) StartPos() *Grid {
	return this.start
}

func (this *World) TargetPos() *Grid {
	return this.target
}

//打印地图
func (this *World) Print() {
	for i := ROW - 1; i >= 0; i-- {
		for j := 0; j < COL; j++ {
			grid := this.grids[i][j]
			fmt.Printf("%v", grid.S)
		}
		fmt.Println("")
	}
}

//打印格子坐标
func (this *World) PrintInfo() {
	for i := ROW - 1; i >= 0; i-- {
		for j := 0; j < COL; j++ {
			grid := this.grids[i][j]
			fmt.Printf("x:%v,y:%v,s:%v ", grid.X, grid.Y, grid.S)
		}
		fmt.Println("")
	}
}

//更新距离目标点距离
func (this *World) UpdateH(g *Grid) {
	xd := int(math.Abs(float64(g.X - this.target.X)))
	yd := int(math.Abs(float64(g.Y - this.target.Y)))
	g.H = xd + yd
}

//定向的格子, 返回从from向to方向的格子, 返回值不包括from和to
func (this *World) Direct(from, to *Grid) []*Grid {
	//from to 必须是相邻点
	xd := int(math.Abs(float64(from.X - to.X)))
	yd := int(math.Abs(float64(from.Y - to.Y)))
	d := xd + yd
	if d != 1 && d != 2 {
		println("求方向的格子必须相邻")
		return nil
	}

	ret := []*Grid{}

	ox := from.X - to.X
	oy := from.Y - to.Y
	for i := 1; i < ROW; i++ {
		x := to.X - ox*i
		y := to.Y - oy*i
		g := this.getGridByPox(x, y)
		if g == nil {
			break
		}
		ret = append(ret, g)
	}

	return ret
}

//寻路
func (this *World) Find() {

}
