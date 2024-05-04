package game

import (
	"strconv"
	"strings"

	"github.com/ikaroly/gobot/pkg/pieces"
)

const StartPos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const Numbers = "0123456789"

type Board struct {
	pieces []pieces.Piece
	moves uint
}

type Chess interface {
	is_checkmate() int
	is_stalemate() bool
}

func NewPosition(fen string) Board{
	var new_board = new(Board)
	if fen == "startpos"{
		fen = StartPos
	}
	
	
	var split_fen = strings.Split(fen, " ")
	var setup = split_fen[0]

	var i_mod = 0
	var new_pieces = []pieces.Piece{}
	for i, char := range setup {
		if char == '/'{
			i_mod -= 1
			continue
		}else if strings.Contains(Numbers, string(char)){
			var diff, _ = strconv.Atoi(string(char))
			i_mod += diff - 1
		}else{
			var new_piece = pieces.NewPiece(char, float64(i + i_mod))
			new_pieces = append(new_pieces, new_piece)
		}
	}
	new_board.pieces = new_pieces

	return *new_board
}