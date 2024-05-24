package search

import (
	"github.com/ikaroly/gobot/pkg/game"
)

func GetBoardMoves(b game.Board) []game.Board{
	var result []game.Board

	for i, piece := range b.Pieces{
		if piece.Color == -b.SideToMove{
			continue
		}

		result = append(result, GetPieceMoves(b, i)...)
	}

	return result
}

func GetPieceMoves(b game.Board, piece_i int) []game.Board{
	var result []game.Board
	var piece = b.Pieces[piece_i]
	var moves = piece.GetMoves(b.BoardCombined, b.EnPassantTarget)

	for _, move := range moves {
		result = append(result, game.MoveBits(b, move))
	}

	return result
}
