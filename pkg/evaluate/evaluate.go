package evaluate

import "github.com/ikaroly/gobot/pkg/game"
// import "github.com/ikaroly/gobot/pkg/pieces"

func Evaluate(b game.Board) int{
	return GetMaterial(b) * 1000
}

func GetMaterial(b game.Board) int{
	var material = 0
	for _, piece := range b.Pieces {
		material += int(piece.Color) * piece.GetValue()
	}
	return material
}