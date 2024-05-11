package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

var BishopMoveVectors = [...]bitboard.Vector{
	{X: 1,Y: 1},
	{X: 1,Y: -1},
	{X: -1,Y: 1},
	{X: -1,Y: -1},
}

type Bishop interface {
	GetBishopMoves()
}

func (p Piece) GetBishopMoves(boardCombined bitboard.CombinedBoard) []bitboard.Move{
	var result []bitboard.Move

	for _, direction := range BishopMoveVectors{
		result = append(result, bitboard.RayCastMovement(p.Position, p.Color, boardCombined, direction)...)
	}
	return result
}