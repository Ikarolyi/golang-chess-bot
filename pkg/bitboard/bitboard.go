package bitboard

import (
	"fmt"
	"math"
)

const files = "ABCDEFGH"

func GetRank(p uint64) uint{
	return uint(math.Floor(float64(GetPlace(p))/8))
}

func GetFile(p uint64) uint{
	return GetPlace(p) % 8
}

func GetPlace(p uint64) uint {
	return uint(math.Log2(float64(p)))
}

func ToString(p uint64) string {
	return string(files[GetRank(p) + 1]) + fmt.Sprint(GetRank(p))
}