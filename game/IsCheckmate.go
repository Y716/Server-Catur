package game

import "github.com/Y716/Server-Catur/board"


func IsCheckmate(Board *[8][8]board.Piece, colorFlag bool) bool{
	kingColor := board.NoPieceColor
	if colorFlag{
		kingColor = board.White
	} else{
		kingColor = board.Black
	}

	for i := 0; i < len(Board); i++ {
		for j := 0; j < len(Board); j++ {
			if Board[i][j].PieceColor == kingColor{

				for k := 0; k < len(Board); k++ {
					for l := 0; l < len(Board); l++ {
						//TODO
						fromFile := 
						if !isValidMove(Board, fromFile, toFile, fromRank, toRank) {
							return false
						}
						simulationBoard := 
						IsKingInCheckAfterMove()
					}
				}
			}  
		}
	}


	return false
}

