package main

import (
	"fmt"
	"testing"
)

const worldTest string =
`
S...T...
........
........
B.BBB...
........
.......B
........
........
`

func TestNeighbors(t *testing.T) {
	world := &World{}
	ok := world.LoadWorld(worldTest)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	ns := world.Neighbors(world.start.X, world.start.Y)
	for _, v := range ns {
		v.S =  "N"
	}

	world.Print()
	//world.PrintPox()
}

