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

func (g *Grid) isBlock() bool {
	return g.S == StateBlock
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
	for _, v := range DirectOffset {
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
func (this *World) Direct(from, to *Grid) (ret []*Grid) {
	//from to 必须是相邻点
	xd := int(math.Abs(float64(from.X - to.X)))
	yd := int(math.Abs(float64(from.Y - to.Y)))
	d := xd + yd
	if d != 1 && d != 2 {
		println("求方向的格子必须相邻")
		return nil
	}

	ret = []*Grid{}
	ox := from.X - to.X
	oy := from.Y - to.Y
	for i := 1; i < ROW; i++ {
		x := to.X - ox*i
		y := to.Y - oy*i
		g := this.getGridByPox(x, y)
		if g == nil || g.isBlock() {
			break
		}
		ret = append(ret, g)
	}

	return ret
}

//与目标点直连
func (this *World) Straight(g *Grid) bool {
	if g.X == this.target.X && g.Y == this.target.Y {
		return true
	}

	//纵向直连
	if g.X == this.target.X {
		return this.yStraight(g)
	}

	//横向直连
	if g.Y == this.target.Y {
		return this.xStraight(g)
	}

	//斜线直连
	xd := int(math.Abs(float64(g.X - this.target.X)))
	yd := int(math.Abs(float64(g.Y - this.target.Y)))
	if xd == yd {
		return this.crossStraight(g)
	}

	return false
}

func (this *World) yStraight(g *Grid) bool {

	min, max := g.Y, this.target.Y
	if min > max {
		min, max = max, min
	}

	for i := min + 1; i < max; i++ {
		g1 := this.getGridByPox(g.X, i)
		if g1.isBlock() {
			return false
		}
	}

	return true
}

func (this *World) xStraight(g *Grid) bool {

	min, max := g.X, this.target.X
	if min > max {
		min, max = max, min
	}

	for i := min + 1; i < max; i++ {
		g1 := this.getGridByPox(i, g.Y)
		if g1.isBlock() {
			return false
		}
	}

	return true
}

func (this *World) crossStraight(g *Grid) bool {
	tmpY := []int{}
	startY, endY := g.Y, this.target.Y
	if startY < endY {
		for i := startY + 1; i < endY; i++ {
			tmpY = append(tmpY, i)
		}
	} else {
		for i := startY - 1; i > endY; i-- {
			tmpY = append(tmpY, i)
		}
	}

	tmpX := []int{}
	startX, endX := g.X, this.target.X
	if startX < endX {
		for i := startX + 1; i < endX; i++ {
			tmpX = append(tmpX, i)
		}
	} else {
		for i := startX - 1; i > endX; i-- {
			tmpX = append(tmpX, i)
		}
	}

	for i := 0; i < len(tmpY); i++ {
		g2 := this.getGridByPox(tmpX[i], tmpY[i])
		if g2.isBlock() {
			return false
		}
	}

	return true
}

//寻路
func (this *World) Find() {

}
