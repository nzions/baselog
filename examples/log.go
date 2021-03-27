package main

import "github.com/nzions/baselog"

type MyThing struct {
	Logger baselog.Logger
}

func (x *MyThing) SayHello() {
	x.Logger.Logf("Hello")
}

func main() {

	mt := MyThing{
		Logger: &baselog.STDOutLog{},
	}
	mt.SayHello()
}
