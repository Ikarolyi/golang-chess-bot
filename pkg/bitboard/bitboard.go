package bitboard

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func Encode(in string) uint64 {
	var splitin = strings.Split(in, "")
	var rank, _ = strconv.Atoi(splitin[1])
	var file_ascii = splitin[0][0]
	var file = 0

	
	if file_ascii > 90{
		file = int(file_ascii) - 96 //Uppercase
	}else{
		file = int(file_ascii) - 64 //Lowercase
	}

	return uint64(math.Pow(2, float64(file + rank*8)))
}