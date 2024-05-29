package main

import (
	"github.com/ikaroly/gobot/pkg/bitboard"
	"github.com/ikaroly/gobot/pkg/uci"
)


func main() {
	var engine = new(uci.Engine)


	engine.Init()
	engine.Listen()

	println("a1 is ", bitboard.GetRank(bitboard.Encode("a1")))

}