package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

type Pawn interface {
	GetPawnMoves() []bitboard.Move
}

func (p Piece) GetPawnMoves(boardCombined bitboard.CombinedBoard, enPassantTarget uint64) []bitboard.Move {
	var result []bitboard.Move
	return result
}