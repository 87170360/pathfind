package main

import (
	"fmt"
	"testing"
)

func TestNeighbors(t *testing.T) {
	const worldTest string = `
....T...
........
........
B.BBB...
........
.......B
........
S.......
`
	world := &World{}
	ok := world.LoadWorld(worldTest)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	ns := world.Neighbors(world.start.X, world.start.Y)
	for _, v := range ns {
		v.S = "N"
	}

	world.Print()
	//world.PrintPox()
}

func TestDirect(t *testing.T) {
	const worldTest string = `
........
........
........
........
........
........
.T......
S.......
`
	world := &World{}
	ok := world.LoadWorld(worldTest)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	tmp := world.Direct(world.start, world.target)
	for _, v := range tmp {
		v.S = "D"
	}

	world.Print()
}
