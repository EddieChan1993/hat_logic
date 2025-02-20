package ants

import (
	ants2 "github.com/panjf2000/ants/v2"
)

const poolSize = 100

var AntsGo *ants2.Pool

func InitAnts() {
	var err error
	AntsGo, err = ants2.NewPool(poolSize)
	if err != nil {
		panic(err)
	}
}

func ReleaseAnts() {
	AntsGo.Release()
}
