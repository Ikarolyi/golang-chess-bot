package main

import (
	"github.com/ikaroly/gobot/pkg/uci"
)


func main() {
	var engine = new(uci.Engine)


	engine.Init()
	engine.Listen()

}