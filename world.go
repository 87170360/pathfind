package main

import (
	"fmt"
	"strings"
)

//格子
type Grid struct {
	S string //状态
	X int    //坐标
	Y int    //坐标
	H int    //路径增量H值, 使用曼哈顿距离方法
	G int    //路径增量G值
}

//获取路径增量
func (g *Grid) F() int {
	return g.H + g.G
}

func (g *Grid) isStart() bool {
	return g.S == StateStart
}

func (g *Grid) isTarget() bool {
	return g.S == StateTarget
}

//数组第一维代表1行row
type World struct {
	grids [COL][ROW]*Grid
	start *Grid
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
				X: c,
				Y: j,
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
