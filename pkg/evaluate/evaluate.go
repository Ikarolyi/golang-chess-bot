package evaluate

import "github.com/ikaroly/gobot/pkg/game"
// import "github.com/ikaroly/gobot/pkg/pieces"

func Evaluate(b game.Board) int{
	return material(b) * 1000
}

func material(b game.Board) int{
	var material = 0
	for _, piece := range b.Pieces {
		material += piece.Color * piece.GetValue()
	}
	return material
}