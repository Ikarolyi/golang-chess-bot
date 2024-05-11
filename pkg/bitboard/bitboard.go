package bitboard

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const files = "abcdefgh"

const FileA uint64 = 0b1000000010000000100000001000000010000000100000001000000010000000
const Rank8 uint64 = 0b11111111

const Diagonal uint64 = 		0b1000000001000000001000000001000000001000000001000000001000000001
const AntiDiagonal uint64 = 0b0000000100000010000001000000100000010000001000000100000010000000

type Move struct {
	From uint64
	To uint64
}

// if + {pos << by} if - {pos >> |by|}
func BitShiftPos(pos uint64, by int) uint64{
	if by > 0{
		return pos << by
	}else{
		return pos >> by
	}
}

func GetRank(p uint64) uint{
	return uint(8 - math.Floor(float64(GetPlace(p))/8))
}

func GetFile(p uint64) uint{
	return (GetPlace(p) % 8)
}

func GetPlace(p uint64) uint {
	return uint(math.Log2(float64(p)))
}

func ToString(p uint64) string {
	return string(files[GetFile(p)]) + fmt.Sprint(GetRank(p))
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

func IsSquareEmpty(square uint64, boardCombined uint64) bool{
	return square == boardCombined & square
}

func GetDistanceFromEdge(square uint64, direction Vector) int{
	var file = GetFile(square)
	var rank = GetRank(square)

	var left_distance = file
	var right_distance = 8 - file

	var upper_distance = rank
	var lower_distance = 8 - 	rank

	var result = 10.0

	if direction.X < 0 {
		result = math.Min(result, float64(left_distance))
	}else if direction.X > 0{
		result = math.Min(result, float64(right_distance))
	}
	if direction.Y < 0{
		result = math.Min(result, float64(lower_distance))
	}else if direction.Y > 0{
		result = math.Min(result, float64(upper_distance))
	}

	return int(result)
}

func RayCastMovement(square uint64, color int, boardCombined CombinedBoard, direction Vector) []Move {
	var distance = GetDistanceFromEdge(square, direction)
	var result []Move

	for i := 1; i <= distance; i++{
		var pos = Translate(square, direction.Multiply(i))
		var occupied = IsSquareEmpty(pos, boardCombined.GetTrueCombined())
		var occupiedByFriend = IsSquareEmpty(pos, boardCombined.GetColor(color))
		if !occupiedByFriend{
			result = append(result, Move{square, pos})
		}

		if occupied{
			break
		}
	}
	return result
}