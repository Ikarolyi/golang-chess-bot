package bitboard

type Vector struct {
	X, Y int
}

var RookMoveVectors = [...]Vector{
	{0,1},
	{1,0},
	{0,-1},
	{-1,0},
}



func Translate(point uint64, by Vector) uint64{
	point = BitShiftPos(point, by.X)
	point = BitShiftPos(point, by.Y * 8)
	return point
}

func (v Vector) Multiply(by int) Vector{
	return Vector{v.X * by, v.Y * by}
}