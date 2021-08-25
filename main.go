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

func main() {
	//fmt.Printf("%v", Conf.World)
	for _, v := range Conf.World {
		work(v)
		fmt.Println("=========================================")
	}

}

func work(data string) {
	world := &World{}
	ok := world.LoadWorld(data)
	if !ok {
		return
	}

	world.Print()

	fmt.Println("------------")

	step, ok := world.FindStep()
	fmt.Printf("found :%v\n", ok)

	world.SetPathState(step)
	world.SetStepState(step)
	world.Print()
}
