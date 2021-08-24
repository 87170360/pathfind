package main

import (
	"container/heap"
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
	stand  *Grid
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
				this.stand = grid
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

func (this *World) Stand() *Grid {
	return this.stand
}

func (this *World) Target() *Grid {
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
func (this *World) Straight(start, end *Grid) (path []*Grid, straight bool) {
	if start.X == end.X && start.Y == end.Y {
		return nil, true
	}

	//纵向直连
	if start.X == end.X {
		return this.yStraight(start, end)
	}

	//横向直连
	if start.Y == end.Y {
		return this.xStraight(start, end)
	}

	//斜线直连
	xd := int(math.Abs(float64(start.X - end.X)))
	yd := int(math.Abs(float64(start.Y - end.Y)))
	if xd == yd {
		return this.crossStraight(start, end)
	}

	return
}

func (this *World) yStraight(start, end *Grid) (path []*Grid, straight bool) {
	startY, endY := start.Y, end.Y
	if startY < endY {
		for i := startY + 1; i < endY; i++ {
			g1 := this.getGridByPox(start.X, i)
			if g1.isBlock() {
				straight = false
				return
			}
			path = append(path, g1)
		}
	} else {
		for i := startY - 1; i > endY; i-- {
			g1 := this.getGridByPox(start.X, i)
			if g1.isBlock() {
				straight = false
				return
			}
			path = append(path, g1)
		}
	}

	straight = true
	return
}

func (this *World) xStraight(start, end *Grid) (path []*Grid, straight bool) {
	startX, endX := start.X, end.X
	if startX < endX {
		for i := startX + 1; i < endX; i++ {
			g1 := this.getGridByPox(i, start.Y)
			if g1.isBlock() {
				straight = false
				return
			}
			path = append(path, g1)
		}
	} else {
		for i := startX - 1; i > endX; i-- {
			g1 := this.getGridByPox(i, start.Y)
			if g1.isBlock() {
				straight = false
				return
			}
			path = append(path, g1)
		}
	}

	straight = true
	return
}

func (this *World) crossStraight(start, end *Grid) (path []*Grid, straight bool) {
	tmpY := []int{}
	startY, endY := start.Y, end.Y
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
	startX, endX := start.X, end.X
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
			straight = false
			return
		}
		path = append(path, g2)
	}

	straight = true
	return
}

func (this *World) CreatePQ(gs []*Grid) *PriorityQueue {
	pg := make(PriorityQueue, len(gs))
	i := 0
	for _, v := range gs {
		pg[i] = &Item{
			value:    v,
			priority: v.H,
			index:    i,
		}
		i++
	}
	heap.Init(&pg)

	return &pg
}

func (this *World) PQPop(pq *PriorityQueue) *Grid {
	if pq.Len() <= 0 {
		return nil
	}

	item := heap.Pop(pq).(*Item)
	return item.value.(*Grid)
}

//寻路， 最多只选择2步可达路径, step是每一步的点， path路径上的格子
func (this *World) Find() (step, path []*Grid, find bool) {
	//开始点直连接目标
	if _, ok := this.Straight(this.stand, this.target); ok {
		find = true
		return
	}

	//取起点相邻点
	n := this.Neighbors(this.stand.X, this.stand.Y)
	//计算相邻点权重
	for _, v := range n {
		this.UpdateH(v)
	}

	//根据权重生成优先队列
	pq := this.CreatePQ(n)
	for pq.Len() > 0 && !find {
		//从队列中取最优先格子
		g := this.PQPop(pq)

		//判断格子是否直连目标点
		if p, ok := this.Straight(g, this.target); ok {
			step = append(step, g)
			path = append(path, p...)
			//直连目标，结束寻路
			break
		}

		//非直接可达, 获取起点为原点，优先格子为方向，射线方向的通路格子
		d := this.Direct(this.stand, g)
		//遍历通路格子是否直达目标点
		for _, v := range d {
			if p, ok := this.Straight(v, this.target); ok {
				if p2, ok2 := this.Straight(g, v); ok2 {
					step = append(step, g)
					path = append(path, g)
					path = append(path, v)
					path = append(path, p2...)
				}

				step = append(step, v)
				path = append(path, p...)

				//直连目标，结束寻路
				find = true
				break
			}
		}
	}

	//不直连，标记最近位置
	return
}
