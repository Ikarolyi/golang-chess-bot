package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ikaroly/gobot/pkg/bitboard"
	"github.com/ikaroly/gobot/pkg/pieces"
	"golang.design/x/reflect"
)

const StartPos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
// const StartPos = "8/7P/8/8/8/8/8/7P w KQkq - 0 1"
const Numbers = "0123456789"

var CastlingMoves = [...]bitboard.Move{
	{From: 16, To: 4}, // e1c1
	{From: 16, To: 64}, // e1g1
	{From: 1152921504606846976, To: 288230376151711744}, // e8c8
	{From: 1152921504606846976, To: 4611686018427387904}, // e8g8
}

var RookCastlingMoves = [...]bitboard.Move{
	{From: 1, To: 8}, // a1d1
	{From: 128, To: 32}, // h1f1
	{From: 72057594037927936, To: 576460752303423488}, // a8d8
	{From: 9223372036854775808, To: 2305843009213693952}, // h8f8
}


type Board struct {
	Pieces []pieces.Piece
	Moves int
	SideToMove int8
	Castling byte //0b0000KQkq
	EnPassantTarget uint64
	HalfmoveCounter int
	FullmoveCounter int
	BoardCombined bitboard.CombinedBoard
	LastMove bitboard.Move
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
	var setup = strings.Trim(split_fen[0], "/")
	
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
			var new_piece = pieces.NewPiece(char, i+i_mod) // Reversed to match the board notation
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

func GetCastledSide(move bitboard.Move) int {
	var castledSide int = -1

	for i, castlingMove := range CastlingMoves {
		if castlingMove == move{
			castledSide = i
			break
		}
	}

	return castledSide
}

func MoveBits(b Board, move bitboard.Move) Board{
	var new_board = reflect.DeepCopy(b)
	piece_to_move_i, piece_to_move := b.GetPieceOnPos(move.From)
	captured_piece_i, _ := b.GetPieceOnPos(move.To)

	if piece_to_move_i == pieces.EMPTYINDEX{
		return b // Invalid move
	}

	// Castling
	if piece_to_move.Class == pieces.KING{
		castlingSide := GetCastledSide(move)

		// Castling happened
		if castlingSide != -1{

			println("Google castling")
			side_to_move := b.SideToMove // Cache side to move because moving the rook will change it
			
			// Move the rook
			var rookMove = RookCastlingMoves[castlingSide]
			new_board = MoveBits(new_board, rookMove)

			new_board.SideToMove = side_to_move // reset side to move
		}
	}

	// TODO optimization

	// Actually make the move
	new_board.Pieces[piece_to_move_i].Position = move.To
	

	// Remove the captured piece, if there's one
	if captured_piece_i != pieces.EMPTYINDEX{
		new_board.Pieces = GetSetWithoutPieceI(new_board.Pieces, captured_piece_i)
	}
	
	new_board.BoardCombined = new_board.GetBoardCombined()
	new_board.SideToMove = -new_board.SideToMove
	new_board.EnPassantTarget = move.NewEnPassantTarget

	new_board.LastMove = move
	
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
	return append(piece_set[:index], piece_set[index+1:]...)
}

func (b *Board) MoveAll(moves []string){
	for _, move := range moves{
		if len(move) == 4{
			b.Move(move)
		}
	}
}

func (b Board) ToString() string{
	var retVal = ""
	
	for _, p := range b.Pieces{
		retVal += p.ToString() + " | "
	}
	
	return retVal
}

func (b Board) ExportFEN() string {
	var result = ""
	
	var emptyFor = 0
	for sqI := 0; sqI < 63; sqI++{
		// Flip files (fen thing)
		var bitSquare uint64 = 1 << (56 - (int(sqI/8) * 8) + (sqI % 8))


		var pI, piece = b.GetPieceOnPos(bitSquare)
		var newLine = ((sqI % 8) == 0) && sqI != 0

		if newLine{
			if emptyFor != 0{
				result += fmt.Sprint(emptyFor)
				emptyFor = 0
			}
			result += "/"
		}

		if pI == pieces.EMPTYINDEX{
			emptyFor += 1
		}else{
			if emptyFor != 0{
				result += fmt.Sprint(emptyFor)
				emptyFor = 0
			}
			result += piece.ToString()
		}
	}

	if emptyFor != 0{
		result += fmt.Sprint(emptyFor)
		emptyFor = 0
	}

	return result
}