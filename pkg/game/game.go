package game

import (
	"strconv"
	"strings"

	"github.com/ikaroly/gobot/pkg/bitboard"
	"github.com/ikaroly/gobot/pkg/pieces"
	"golang.design/x/reflect"
)

const StartPos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
// const StartPos = "8/7P/8/8/8/8/8/7P w KQkq - 0 1"
const Numbers = "0123456789"


type Board struct {
	Pieces []pieces.Piece
	Moves int
	SideToMove int8
	Castling byte //0b0000KQkq
	EnPassantTarget uint64
	HalfmoveCounter int
	FullmoveCounter int
	BoardCombined bitboard.CombinedBoard
}

type Chess interface {
	is_checkmate() int
	is_stalemate() bool
	MoveAll()
	Move()
}


func NewPosition(fen string) Board{
	var new_board = new(Board)
	if fen == "startpos"{
		fen = StartPos
	}
	
	
	var split_fen = strings.Split(fen, " ")
	var setup = split_fen[0]
	
	// Set up position
	var i_mod = 0
	var new_pieces = []pieces.Piece{}
	for i, char := range setup {
		if char == '/'{
			i_mod -= 1
			continue
		}else if strings.ContainsRune(Numbers, char){
			var diff, _ = strconv.Atoi(string(char))
			i_mod += diff - 1
		}else{
			var new_piece = pieces.NewPiece(char, float64(63 - (i + i_mod))) // Reversed to match the board notation
			new_pieces = append(new_pieces, new_piece)
			
		}
	}
	new_board.Pieces = new_pieces
	
	// Set up color
	if split_fen[1] == "w"{
	new_board.SideToMove = pieces.WHITE
	}else{
		new_board.SideToMove = pieces.BLACK
	}
	
	//Set up castling
	var castling_fen = split_fen[2]
	new_board.Castling = byte(0b00000000)
	if strings.Contains(castling_fen, "K"){
		new_board.Castling |= byte(0b00001000)
	}
	if strings.Contains(castling_fen, "Q"){
		new_board.Castling |= byte(0b00000100)
	}
	if strings.Contains(castling_fen, "k"){
		new_board.Castling |= byte(0b00000010)
	}
	if strings.Contains(castling_fen, "q"){
		new_board.Castling |= byte(0b00000001)
	}
	
	//Set up en passant
	new_board.EnPassantTarget = 0
	if split_fen[3] != "-"{
		new_board.EnPassantTarget = bitboard.Encode(split_fen[3])
	}
	
	//Counters
	new_board.HalfmoveCounter, _ = strconv.Atoi(split_fen[4])
	new_board.FullmoveCounter, _ = strconv.Atoi(split_fen[5])
	
	return *new_board
}

func (b Board) GetPieceOnPos(position uint64) (int, pieces.Piece){
	// Faster than using bitboard.IsSquareEmpty()

	// No exception for multiple pieces on the same square
	for index, piece := range b.Pieces{

		if piece.Position == position{
			return index, piece
		}
	}

	return pieces.EMPTYINDEX, pieces.NewPiece('P', 1)
}


func MoveBits(b Board, move bitboard.Move) Board{
	var new_board = reflect.DeepCopy(b)
	piece_to_move_i, _ := b.GetPieceOnPos(move.From)
	captured_piece_i, _ := b.GetPieceOnPos(move.To)

	// TODO optimization

	// for i, piece := range new_board.Pieces{
	// 	if piece.Position == move.From{
	// 		piece_to_move_i = i
	// 		println("f",bitboard.ToString(piece.Position))
	// 	}else if piece.Position == move.To{
	// 		captured_piece_i = i
	// 		println("t", bitboard.ToString(piece.Position))
	// 	}
	// }

	// Actually make the move
	new_board.Pieces[piece_to_move_i].Position = move.To
	

	// Remove the captured piece, if there's one
	if captured_piece_i != pieces.EMPTYINDEX{
		new_board.Pieces = GetSetWithoutPieceI(new_board.Pieces, captured_piece_i)
	}
	
	new_board.BoardCombined = new_board.GetBoardCombined()
	new_board.SideToMove = -new_board.SideToMove
	
	return new_board
}

func (b Board) GetBoardCombined() bitboard.CombinedBoard {
	var white uint64 = 0
	var black uint64 = 0
	for _, piece := range b.Pieces{
		if piece.Color == pieces.WHITE{
			white |= piece.Position
			}else{
				black |= piece.Position
			}
		}
		return bitboard.CombinedBoard{White: white, Black: black}
	}
	
	func (b *Board) Move(moveFromTo string) {
		var from = bitboard.Encode(moveFromTo[0:2])
		var to = bitboard.Encode(moveFromTo[2:4])
		
		*b = MoveBits(*b, bitboard.Move{From: from, To: to})
		b.Moves += 1
	}
	
	func GetSetWithoutPieceI(piece_set []pieces.Piece, index int) []pieces.Piece{
		// b.Pieces = append(b.Pieces[:index], b.Pieces[index+1:]...)
		return append(piece_set[:index], piece_set[index+1:]...)
	}
	
	func (b *Board) MoveAll(moves []string){
		for _, move := range moves{
			b.Move(move)
		}
	}
	
	func (b Board) ToString() string{
		var retVal = ""
		
		for _, p := range b.Pieces{
			retVal += p.ToString() + " | "
		}
		
		return retVal
	}