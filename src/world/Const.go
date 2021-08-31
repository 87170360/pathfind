package world

const (
	ROW int = 8
	COL int = 8
)

const (
	StateStart  string = "S" //开始点
	StateTarget string = "T" //目标点
	StateFlat   string = "." //平路
	StateBlock  string = "B" //障碍
	StatePath   string = "*" //路径点
)

//相邻格子偏移量 上，下，左，右，上左，上右，下左，下右
var DirectOffset = [8][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {1, 1}, {-1, -1}, {1, -1}}
