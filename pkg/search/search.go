package search

import (
	"github.com/ikaroly/gobot/pkg/bitboard"
	"github.com/ikaroly/gobot/pkg/evaluate"
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

func SearchDepth(b game.Board, master bool, depth uint) (bitboard.Move, int){
	if depth <= 0{
		return bitboard.Move{}, evaluate.Evaluate(b)
	}

	all_moves := GetBoardMoves(b)
	
	var best_score int
	var best_move bitboard.Move

	for i, move := range all_moves{
		_, eval_score := SearchDepth(move, false, depth - 1)
		if i == 0{
			best_score = eval_score
		}else{
			if b.SideToMove == 1 {
				best_score = max(eval_score, best_score)
			}else{
				best_score = min(eval_score, best_score)
			}
		}
		if master{
			if best_score == eval_score{
				best_move = move.LastMove
			}
		}
	}

	return best_move, best_score
}

func MasterSearch(b game.Board, depth uint) (bitboard.Move, int){
	best_move, best_score := SearchDepth(b, true, depth)
	return best_move, best_score
}