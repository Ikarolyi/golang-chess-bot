package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

var QueenMoveVectors = append(RookMoveVectors[:], BishopMoveVectors[:]...)

type Queen interface {
	GetQueenMoves() []bitboard.Move
}

func (p Piece) GetQueenMoves(boardCombined bitboard.CombinedBoard) []bitboard.Move{
	var result []bitboard.Move

	for _, direction := range QueenMoveVectors{
		result = append(result, bitboard.RayCastMovement(p.Position, p.Color, boardCombined, direction, bitboard.UNLIMITED)...)
	}
	return result
}