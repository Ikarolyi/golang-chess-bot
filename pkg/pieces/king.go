package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

var KingMoveVectors = [...]bitboard.Vector{
	{X: 0,Y: 1},
	{X: -1,Y: 0},
	{X: 0,Y: -1},
	{X: 1,Y: 0},

	{X: 1,Y: 1},
	{X: -1,Y: 1},
	{X: -1,Y: -1},
	{X: 1,Y: -1},
}

type King interface {
	GetKingMoves() []bitboard.Move
}

func (p Piece) GetKingMoves(boardCombined bitboard.CombinedBoard) []bitboard.Move{
	var result []bitboard.Move

	for _, direction := range KingMoveVectors{
		result = append(result, bitboard.RayCastMovement(p.Position, p.Color, boardCombined, direction, 1, false)...)
	}
	return result
}