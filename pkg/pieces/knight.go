package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

var KnightMoveVectors = [...]bitboard.Vector{
	{X: 1,Y: 2},
	{X: -1,Y: 2},
	{X: 2,Y: 1},
	{X: -2,Y: 1},
	{X: 1,Y: -2},
	{X: -1,Y: -2},
	{X: 2,Y: -1},
	{X: -2,Y: -1},
}

type Knight interface {
	GetKnightMoves() []bitboard.Move
}

func (p Piece) GetKnightMoves(boardCombined bitboard.CombinedBoard) []bitboard.Move{
	var result []bitboard.Move

	for _, direction := range KnightMoveVectors{
		result = append(result, bitboard.RayCastMovement(p.Position, p.Color, boardCombined, direction, 1, false)...)
	}

	return result
}