package main

import (
	"fmt"
	"testing"
)

func TestWorld_Neighbors(t *testing.T) {
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

func TestWorld_Direct(t *testing.T) {
	const worldTest string = `
........
........
........
....B...
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

func TestWorld_XStraight(t *testing.T) {
	const worldTest string = `
........
........
........
........
........
........
........
TB.....S
`
	world := &World{}
	ok := world.LoadWorld(worldTest)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	b := world.Straight(world.start)
	//output
	fmt.Println(b)
}

func TestWorld_YStraight(t *testing.T) {
	const worldTest string = `
T.......
........
........
........
........
........
B.......
S.......
`
	world := &World{}
	ok := world.LoadWorld(worldTest)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	b := world.Straight(world.start)
	//output
	fmt.Println(b)
}

func TestWorld_CrossStraight(t *testing.T) {
	const worldTest string = `
.......S
........
........
........
........
........
........
T.......
`
	world := &World{}
	ok := world.LoadWorld(worldTest)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	b := world.Straight(world.start)
	//output
	fmt.Println(b)
}
