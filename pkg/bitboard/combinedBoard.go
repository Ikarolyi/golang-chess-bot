package bitboard

type CombinedBoard struct {
	White uint64
	Black uint64
}


func (b CombinedBoard) GetColor(color int8) uint64{
	//TODO a better way to implement combined board, bitboard package shouldn't have colors
	const combinedBoardColorWhite int8 = 1
	if color == combinedBoardColorWhite{
		return b.White
	}else{
		return b.Black
	}
}

func (b CombinedBoard) GetTrueCombined() uint64{
	return b.White | b.Black
}