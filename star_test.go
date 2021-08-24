package main

import (
	"fmt"
	"strconv"
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

	ns := world.Neighbors(world.stand.X, world.stand.Y)
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

	tmp := world.Direct(world.stand, world.target)
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
S......T
`
	world := &World{}
	ok := world.LoadWorld(worldTest)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	path, ok := world.Straight(world.stand, world.target)
	//output
	fmt.Printf("straight :%v\n", ok)
	for _, v := range path {
		v.S = StatePath
	}
	world.Print()
}

func TestWorld_YStraight(t *testing.T) {
	const worldTest string = `
T.......
........
........
B.......
........
........
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

	path, ok := world.Straight(world.stand, world.target)
	//output
	fmt.Printf("straight :%v\n", ok)
	for _, v := range path {
		v.S = StatePath
	}
	world.Print()
}

func TestWorld_CrossStraight(t *testing.T) {
	const worldTest string = `
.......T
........
........
........
........
........
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

	path, ok := world.Straight(world.stand, world.target)
	//output
	fmt.Printf("straight :%v\n", ok)
	for _, v := range path {
		v.S = StatePath
	}
	world.Print()
}

func TestWorld_Find(t *testing.T) {
	const worldTest string = `
.T......
........
........
B.......
........
........
........
S.......
`
	v := worldTest
	fmt.Println(v)
	world := &World{}
	ok := world.LoadWorld(worldTest)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	step, path, ok := world.Find()
	fmt.Printf("found :%v\n", ok)

	for _, v := range path {
		v.S = StatePath
	}

	for i, v := range step {
		v.S = strconv.Itoa(i+1)
	}

	world.Print()
}
