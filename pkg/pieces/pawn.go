package pieces

import "github.com/ikaroly/gobot/pkg/bitboard"

type Pawn interface {
	GetPawnMoves() []bitboard.Move
}

var pawnForward = bitboard.Vector{X: 0, Y: -1}
var pawnKickVectors = [...]bitboard.Vector{
	{X: 1, Y: 1},
	{X: -1, Y: 1},
}

func (p Piece) pawnKick(boardCombined bitboard.CombinedBoard, enPassantTarget uint64) []bitboard.Move{
	var result []bitboard.Move

	if p.Color == WHITE{
		boardCombined.Black |= enPassantTarget
	}else{
		boardCombined.White |= enPassantTarget
	}

	for _, kickMove := range pawnKickVectors{
		targetSquare := bitboard.Translate(p.Position, kickMove.Multiply(int(p.Color)))
		if !bitboard.IsSquareEmpty(targetSquare, boardCombined.GetColor(p.Color)){
			result = append(result, bitboard.Move{From: p.Position, To: targetSquare})
		}
	}

	return result
}

func (p Piece) GetPawnMoves(boardCombined bitboard.CombinedBoard, enPassantTarget uint64) []bitboard.Move {
	var result []bitboard.Move

	var push_len = 1
	var on_default_place = false

	if p.Color == WHITE{
		on_default_place = (bitboard.GetRank(p.Position) == 2)
	}else{
		on_default_place = (bitboard.GetRank(p.Position) == 7)
	}

	if on_default_place{
		push_len = 2
	}

	result = append(result, bitboard.RayCastPawn(p.Position, p.Color, boardCombined, pawnForward.Multiply(int(p.Color)), push_len)...)

	// result = append(result, p.pawnKick(boardCombined, enPassantTarget)...)


	return result
}