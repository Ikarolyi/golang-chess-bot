package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

var RookMoveVectors = [...]bitboard.Vector{
	{X: 0,Y: 1},
	{X: 1,Y: 0},
	{X: 0,Y: -1},
	{X: -1,Y: 0},
}

type Rook interface {
	GetRookMoves() []bitboard.Move
}

func (p Piece) GetRookMoves(boardCombined bitboard.CombinedBoard) []bitboard.Move{
	var result []bitboard.Move

	for _, direction := range RookMoveVectors{
		result = append(result, bitboard.RayCastMovement(p.Position, p.Color, boardCombined, direction, bitboard.UNLIMITED)...)
	}
	return result
}