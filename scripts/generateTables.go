package scripts

import "github.com/ikaroly/gobot/pkg/bitboard"

func GenerateCastlingSkippedSquares(){
	println("--Squares skipped by castling--")

	b1 := bitboard.Encode("b1")
	c1 := bitboard.Encode("c1")
	d1 := bitboard.Encode("d1")
	println(b1 | c1 | d1)

	f1 := bitboard.Encode("f1")
	g1 := bitboard.Encode("g1")
	println(f1 | g1)

	b8 := bitboard.Encode("b8")
	c8 := bitboard.Encode("c8")
	d8 := bitboard.Encode("d8")
	println(b8 | c8| d8)

	f8 := bitboard.Encode("f8")
	g8 := bitboard.Encode("g8")
	println(f8 | g8)

	println("--End--")
}