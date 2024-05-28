/*
A8 ^ 2#3 E8
	 | ###
	 | ###
A1 +-0#1->E1

    ^ The numbers are in the ascending order of the uint64 bitboard exponents
*/
package bitboard

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)


const files = "abcdefgh"

//Move{From, To}
type Move struct {
	From uint64
	To uint64
	NewEnPassantTarget uint64
}

// if + {pos << by} if - {pos >> |by|}
func BitShiftPos(pos uint64, by int) uint64{
	if by > 0{
		return pos << by
	}else{
		return pos >> -by
	}
}

func GetRank(p uint64) uint{
	return uint(math.Floor(float64(GetPlace(p))/8)) + 1
}

func GetFile(p uint64) uint{
	return (GetPlace(p) % 8)
}

func GetPlace(p uint64) uint {
	return uint(math.Log2(float64(p)))
}

func MakeSquare(file int, rank int) uint64{
	return 1 << (file + (8*(rank-1)))
}

func Decode(p uint64) string {
	return string(files[GetFile(p)]) + fmt.Sprint(GetRank(p))
}

func Encode(in string) uint64 {
	var splitin = strings.Split(in, "")
	var rank, _ = strconv.Atoi(splitin[1])

	var file_ascii = splitin[0][0]
	var file = 0

	
	if file_ascii > 90{
		file = int(file_ascii) - 97 //Uppercase 96
	}else{
		file = int(file_ascii) - 65 //Lowercase 64
	}
	_ = rank

	return MakeSquare(file, rank)
}

func IsSquareEmpty(square uint64, boardCombined uint64) bool{
	return square != boardCombined & square
}

func GetDistanceFromEdge(square uint64, direction Vector) int{
	var file = GetFile(square)
	var rank = GetRank(square)

	var left_distance = file
	var right_distance = 8 - file

	var upper_distance = rank
	var lower_distance = 8 - rank

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

const UNLIMITED = 0

//UNLIMITED is a valid limit
func RayCastMovement(square uint64, color int8, boardCombined CombinedBoard, direction Vector, limit int, force_capture bool) []Move {
	var distance = min(GetDistanceFromEdge(square, direction), limit)
	var result []Move

	for i := 1; i <= distance; i++{
		var pos = Translate(square, direction.Multiply(i))
		var occupied = !IsSquareEmpty(pos, boardCombined.GetTrueCombined())
		var occupiedByFriend = !IsSquareEmpty(pos, boardCombined.GetColor(color))
		if !occupiedByFriend{
			if !force_capture{
				result = append(result, Move{From: square, To: pos})
			}else{
				break
			}
		}

		if occupied{
			if force_capture{
				result = append(result, Move{From: square, To: pos})
			}
			break
		}
	}
	return result
}