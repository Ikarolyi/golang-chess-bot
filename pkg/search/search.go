package search

import (
	"github.com/ikaroly/gobot/pkg/bitboard"
	"github.com/ikaroly/gobot/pkg/evaluate"
	"github.com/ikaroly/gobot/pkg/game"
	"github.com/ikaroly/gobot/pkg/pieces"
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

func SearchDepth(b game.Board, master bool, depth uint) int{
	if depth == 0{
		return evaluate.Evaluate(b)
	}

	all_moves := GetBoardMoves(b)
	
	var best int
	for i, move := range all_moves{
		eval_score := SearchDepth(move, false, depth - 1)
		if i == 0{
			best = eval_score
		}else{
			if b.SideToMove == 1 {
				best = max(eval_score, best)
			}else{
				best = min(eval_score, best)
			}
		}
	}

	return best
}

func MasterSearch(b game.Board, depth uint) int{
	return SearchDepth(b, true, depth)
}

// https://www.chessprogramming.org/Negamax <- Negamax