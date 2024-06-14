package tables

import "github.com/ikaroly/gobot/pkg/bitboard"

const (
	WHITE_LONG_CASTLING = 0
	WHITE_SHORT_CASTLING = 1
	BLACK_LONG_CASTLING = 2
	BLACK_SHORT_CASTLING = 3
)

var CastlingKingMove = [...]bitboard.Move{
	{From: 16, To: 4}, // e1c1
	{From: 16, To: 64}, // e1g1
	{From: 1152921504606846976, To: 288230376151711744}, // e8c8
	{From: 1152921504606846976, To: 4611686018427387904}, // e8g8
}

var CastlingRookMove = [...]bitboard.Move{
	{From: 1, To: 8}, // a1d1
	{From: 128, To: 32}, // h1f1
	{From: 72057594037927936, To: 576460752303423488}, // a8d8
	{From: 9223372036854775808, To: 2305843009213693952}, // h8f8
}


var CastlingSkippedSquares = [...]uint64{
	14, // b1 | c1 | d1
	96, // f1 | g1
	1008806316530991104, // b8 | c8 | d8
	6917529027641081856, // f8 | g8
}

// 0b0000KQkq
var CastlingAbility = [...]byte{
	byte(0b00000100), // Q
	byte(0b00001000), // K
	byte(0b00000001), // q
	byte(0b00000010), // k
}