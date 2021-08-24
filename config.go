package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

const file = "conf.toml"


type conf struct {
	World []string `toml:"World"`
}

var Conf *conf

func init() {
	Conf = &conf{}
	_, err := toml.DecodeFile(file, Conf)
	if err != nil {
		fmt.Errorf("decode %v, err:%s\n", file, err.Error())
		panic(err)
	}

}
