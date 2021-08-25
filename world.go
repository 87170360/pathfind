package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
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
			fmt.Printf("%v ", grid.S)
		}
		fmt.Println("")
	}
}

func (this *World) SetPathState(step []*Grid) {
	start := this.stand
	for _, v := range step {
		p, ok := this.Straight(start, v, true)
		if !ok {
			continue
		}

		for _, v2 := range p {
			v2.S = StatePath
		}

		start = v
	}
}

func (this *World) SetStepState(step []*Grid) {
	for i, v := range step {
		v.S = strconv.Itoa(i+1)
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

//定向的格子, 返回从from向to方向的格子(遇到边界或障碍物为止), 返回值不包括from和to
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

//与目标点直连, path不包括首尾的中间点
func (this *World) Straight(start, end *Grid, retPath bool) (path []*Grid, straight bool) {
	if start.X == end.X && start.Y == end.Y {
		return nil, true
	}

	//纵向直连
	if start.X == end.X {
		return this.yStraight(start, end, retPath)
	}

	//横向直连
	if start.Y == end.Y {
		return this.xStraight(start, end, retPath)
	}

	//斜线直连
	xd := int(math.Abs(float64(start.X - end.X)))
	yd := int(math.Abs(float64(start.Y - end.Y)))
	if xd == yd {
		return this.crossStraight(start, end, retPath)
	}

	return
}

func (this *World) yStraight(start, end *Grid, retPath bool) (path []*Grid, straight bool) {
	startY, endY := start.Y, end.Y
	if startY < endY {
		for i := startY + 1; i < endY; i++ {
			g1 := this.getGridByPox(start.X, i)
			if g1.isBlock() {
				straight = false
				return
			}

			if retPath {
				path = append(path, g1)
			}
		}
	} else {
		for i := startY - 1; i > endY; i-- {
			g1 := this.getGridByPox(start.X, i)
			if g1.isBlock() {
				straight = false
				return
			}

			if retPath {
				path = append(path, g1)
			}
		}
	}

	straight = true
	return
}

func (this *World) xStraight(start, end *Grid, retPath bool) (path []*Grid, straight bool) {
	startX, endX := start.X, end.X
	if startX < endX {
		for i := startX + 1; i < endX; i++ {
			g1 := this.getGridByPox(i, start.Y)
			if g1.isBlock() {
				straight = false
				return
			}
			if retPath {
				path = append(path, g1)
			}
		}
	} else {
		for i := startX - 1; i > endX; i-- {
			g1 := this.getGridByPox(i, start.Y)
			if g1.isBlock() {
				straight = false
				return
			}
			if retPath {
				path = append(path, g1)
			}
		}
	}

	straight = true
	return
}

func (this *World) crossStraight(start, end *Grid, retPath bool) (path []*Grid, straight bool) {
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
		if retPath {
			path = append(path, g2)
		}
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

/*
//寻路， 最多只选择2步可达路径, step是每一步的点， path路径上的格子
func (this *World) Find() (step, path []*Grid, find bool) {
	//起点直连目标点
	if p, ok := this.Straight(this.stand, this.target); ok {
		path = append(path, p...)
		step = append(step, this.target)
		find = true
		return
	}

	//找不到路径下的备选格子
	var tmp []*Grid

	//取起点相邻点
	neighbors := this.Neighbors(this.stand.X, this.stand.Y)
	//计算相邻点权重
	for _, v := range neighbors {
		this.UpdateH(v)
	}
	//把相邻点生成优先队列
	pq := this.CreatePQ(neighbors)
	//遍历8个方向的全部格子是否直达目标
	for pq.Len() > 0 && !find {
		//取权重优先的相邻点
		priorityNeighbor := this.PQPop(pq)

		//判断相邻点本身是否直连目标点
		if p, ok := this.Straight(priorityNeighbor, this.target); ok {
			step = append(step, priorityNeighbor)
			path = append(path, p...)
			//直连目标，结束寻路
			find = true
			break
		}

		//起点为原点相邻点为方向，射线方向的点
		direct := this.Direct(this.stand, priorityNeighbor)
		//射线点是否直达目标点
		for _, v := range direct {
			p, ok := this.Straight(v, this.target)
			if !ok {
				continue
			}

			//相邻点到射线点路径
			if p2, ok2 := this.Straight(priorityNeighbor, v); ok2 {
				path = append(path, priorityNeighbor)
				path = append(path, v)
				path = append(path, p2...)
			}

			step = append(step, v)
			path = append(path, p...)

			find = true
			//结束寻路
			break
		}

		if !find && len(direct) > 0 {
			tmp = append(tmp, direct[len(direct)-1])
		}
	}

	if !find && len(tmp) > 0 {
		//不直连，标记最近位置
		for _, v := range tmp {
			this.UpdateH(v)
		}
		pq2 := this.CreatePQ(tmp)
		backup := this.PQPop(pq2)
		if p, ok := this.Straight(this.stand, backup); ok {
			path = append(path, p...)
		}
		step = append(step, backup)
	}

	if find {
		step = append(step, this.target)
	}
	return
}
 */

func (this *World) FindStep() (step []*Grid, find bool) {
	//起点直连目标点
	if _, ok := this.Straight(this.stand, this.target, false); ok {
		step = append(step, this.target)
		find = true
		return
	}

	//找不到路径下的备选格子
	var tmp []*Grid

	//取起点相邻点
	neighbors := this.Neighbors(this.stand.X, this.stand.Y)
	//计算相邻点权重
	for _, v := range neighbors {
		this.UpdateH(v)
	}
	//把相邻点生成优先队列
	pq := this.CreatePQ(neighbors)
	//遍历8个方向的全部格子是否直达目标
	for pq.Len() > 0 && !find {
		//取权重优先的相邻点
		priorityNeighbor := this.PQPop(pq)

		//判断相邻点本身是否直连目标点
		if _, ok := this.Straight(priorityNeighbor, this.target, false); ok {
			step = append(step, priorityNeighbor)
			//直连目标，结束寻路
			find = true
			break
		}

		//起点为原点相邻点为方向，射线方向的点
		direct := this.Direct(this.stand, priorityNeighbor)
		//射线点是否直达目标点
		for _, v := range direct {
			_, ok := this.Straight(v, this.target, false)
			if !ok {
				continue
			}
			step = append(step, v)
			find = true
			//结束寻路
			break
		}

		if !find && len(direct) > 0 {
			tmp = append(tmp, direct[len(direct)-1])
		}
	}

	//没有路径，从备选点中选择权重高的
	if !find && len(tmp) > 0 {
		for _, v := range tmp {
			this.UpdateH(v)
		}
		pq2 := this.CreatePQ(tmp)
		backup := this.PQPop(pq2)
		step = append(step, backup)
	}

	if find {
		step = append(step, this.target)
	}
	return
}
