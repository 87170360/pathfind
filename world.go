package main

import "fmt"

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


//打印地图
func PrintWorld(world *World) {
	for i := ROW - 1; i >= 0; i-- {
		for j := 0; j < COL; j++ {
			grid := world[i][j]
			fmt.Printf("%v", grid.S)
		}
		fmt.Println("")
	}
}




