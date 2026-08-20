package game

import(
	"github.com/Y716/Server-Catur/board"

)

func isKingInCheck(Board *[8][8]board.Piece, colorFlag bool) bool{
	kingColor := board.NoPieceColor
	if colorFlag{
		kingColor = board.White
	} else{
		kingColor = board.Black
	}
	kingFile := 0
	kingRank := 0
	for i := 0; i < len(Board); i++ {
		for j := 0; j < len(Board); j++ {
			if Board[i][j].PieceType == board.King && Board[i][j].PieceColor == kingColor{
				kingRank = i
				kingFile = j
			}  
		}
	}

	for i := 0; i < len(Board); i++ {
		for j := 0; j < len(Board); j++ {
			if Board[i][j].PieceColor != kingColor && Board[i][j].PieceColor != board.NoPieceColor{
				if isValidMove(Board, j, kingFile, i, kingRank){
					return true
				}
			}  
		}
	}

	return false
}
