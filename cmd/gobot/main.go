package main

// import "fmt"
import "github.com/ikaroly/gobot/pkg/uci"


func main() {
	// println(bitboard.ToString( uint64(math.Pow(2, 8)) ))
	var engine = new(uci.Engine)


	engine.Init()
	engine.Listen()

}