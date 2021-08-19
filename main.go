package main

import (
	"fmt"
)

/*
row:2 |
row:1 |
row:0 |---------
		col:0 col:1 col:2
*/

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


func main() {
	fmt.Println("Hello, A star!")
	world := &World{}
	ok := world.LoadWorld(world1)
	if !ok {
		return
	}
	world.Print()
}
