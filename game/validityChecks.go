package game

import (
	"math"

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

func IsKingInCheckAfterMove(Board [8][8]board.Piece, fromRank int, fromFile int, toRank int, toFile int, colorFlag bool) bool{
	simulationBoard := Board

	movedPiece := simulationBoard[fromRank][fromFile]
	simulationBoard[fromRank][fromFile] = board.Piece{
		PieceType:  board.NoPieceType,
		PieceColor: board.NoPieceColor,
	}
	simulationBoard[toRank][toFile] = movedPiece

	return isKingInCheck(&simulationBoard, colorFlag)
}

func IsCheckmate(Board [8][8]board.Piece, colorFlag bool) bool{
	if !HasLegalMoves(&Board, colorFlag)&& isKingInCheck(&Board, colorFlag){
			return true
	}
	return false

}
func HasLegalMoves(Board *[8][8]board.Piece, colorFlag bool) bool{
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

						fromRank := i 
						fromFile := j 
						toRank := k 
						toFile := l 
						simulationBoard := *Board
						if isValidMove(&simulationBoard, fromFile, toFile, fromRank, toRank) && !IsKingInCheckAfterMove(simulationBoard, fromRank, fromFile, toRank, toFile, colorFlag){
							return true
						}

					}
				}
			}  
		}
	}


	return false
}


func isValidMove(Board *[8][8]board.Piece, fromFile, toFile, fromRank, toRank int) bool {
	pieceType := Board[fromRank][fromFile].PieceType
	pieceColor := Board[fromRank][fromFile].PieceColor


	if toFile == fromFile && toRank == fromRank{
		return false
	}

	switch pieceType {
	case board.Pawn:
		if Board[toRank][toFile].PieceType == board.NoPieceType {
			if fromFile == toFile {
				if fromRank == 1 && pieceColor == board.Black || fromRank == 6 && pieceColor == board.White {
					diff := fromRank - toRank
					if ((diff == -1 || diff == -2) && pieceColor == board.Black) ||
						((diff == 1 || diff == 2) && pieceColor == board.White) {

						if (diff == 2 && Board[fromRank-1][fromFile].PieceType != board.NoPieceType) ||
							(diff == -2 && Board[fromRank+1][fromFile].PieceType != board.NoPieceType) {
							break
						}
						return true
					} else {
						break
					}
				} else {
					diff := fromRank - toRank
					if diff == -1 && pieceColor == board.Black || diff == 1 &&
						pieceColor == board.White {
						return true
					} else {
						break
					}
				}
			} else {
				break
			}
		} else {
			if (toFile == fromFile+1 || toFile == fromFile-1) &&
				(Board[toRank][toFile].PieceColor != pieceColor) {
				diff := fromRank - toRank
				if diff == -1 && pieceColor == board.Black ||
					diff == 1 && pieceColor == board.White {
					return true
				} else {
					break
				}
			}

		}
	case board.Rook:

		if Board[toRank][toFile].PieceColor == pieceColor {
			break
		}

		if (toFile == fromFile && fromRank != toRank) || (toRank == fromRank && fromFile != toFile) {
			if fromRank != toRank {
				if fromRank < toRank {
					for i := fromRank + 1; i < toRank; i++ {
						if Board[i][fromFile].PieceType != board.NoPieceType {
							return false
						}
					}
				} else {
					for i := fromRank - 1; i > toRank; i-- {
						if Board[i][fromFile].PieceType != board.NoPieceType {
							return false
						}
					}
				}

			} else if fromFile != toFile {
				if fromFile < toFile {
					for i := fromFile + 1; i < toFile; i++ {
						if Board[fromRank][i].PieceType != board.NoPieceType {
							return false
						}
					}
				} else {
					for i := fromFile - 1; i > toFile; i-- {
						if Board[fromRank][i].PieceType != board.NoPieceType {
							return false
						}
					}
				}
			}

			return true

		}

	case board.Knight:
		if (math.Abs(float64(toFile-fromFile)) == 2 && math.Abs(float64(toRank-fromRank)) == 1) ||
			(math.Abs(float64(toFile-fromFile)) == 1 && math.Abs(float64(toRank-fromRank)) == 2) {
			if pieceColor != Board[toRank][toFile].PieceColor {
				return true
			}
		}
		return false

	case board.Bishop:

		if Board[toRank][toFile].PieceColor == pieceColor {
			return false
		}

		if math.Abs(float64(toFile-fromFile)) == math.Abs(float64(toRank-fromRank)) {
			diffFile := toFile - fromFile
			diffRank := toRank - fromRank
			f := fromFile
			r := fromRank

			if diffFile > 0 && diffRank > 0 {
				f++
				r++
				for range diffFile - 1 {
					if Board[r][f].PieceType != board.NoPieceType {
						return false
					}
					f++
					r++
				}
			} else if diffFile > 0 && diffRank < 0 {
				f++
				r--
				for range diffFile - 1 {
					if Board[r][f].PieceType != board.NoPieceType {
						return false
					}
					f++
					r--
				}
			} else if diffFile < 0 && diffRank < 0 {
				f--
				r--
				for range diffFile*(-1) - 1 {
					if Board[r][f].PieceType != board.NoPieceType {
						return false
					}
					f--
					r--
				}
			} else if diffFile < 0 && diffRank > 0 {
				f--
				r++
				for range diffRank - 1 {
					if Board[r][f].PieceType != board.NoPieceType {
						return false
					}
					f--
					r++
				}
			}
			return true
		}

	case board.Queen:
		// Queen bisa gerak horizontal, vertikal (File beda tapi Rank sama atau Rank beda File sama (antara fromSquare dan toSquare), dan diagonal (selisih File dan Rank sama)
		// Queen terhalang oleh bidak ke arah toSquarenya.Jika ada bidak, maka invalid.

		if Board[toRank][toFile].PieceColor == pieceColor {
			return false
		}

		if (toFile == fromFile && fromRank != toRank) || (toRank == fromRank && fromFile != toFile) {
			if fromRank != toRank {
				if fromRank < toRank {
					for i := fromRank + 1; i < toRank; i++ {
						if Board[i][fromFile].PieceType != board.NoPieceType {
							return false
						}
					}
				} else {
					for i := fromRank - 1; i > toRank; i-- {
						if Board[i][fromFile].PieceType != board.NoPieceType {
							return false
						}
					}
				}

			} else if fromFile != toFile {
				if fromFile < toFile {
					for i := fromFile + 1; i < toFile; i++ {
						if Board[fromRank][i].PieceType != board.NoPieceType {
							return false
						}
					}
				} else {
					for i := fromFile - 1; i > toFile; i-- {
						if Board[fromRank][i].PieceType != board.NoPieceType {
							return false
						}
					}
				}
			}
			return true
		} else {
			if math.Abs(float64(toFile-fromFile)) == math.Abs(float64(toRank-fromRank)) {
				diffFile := toFile - fromFile
				diffRank := toRank - fromRank
				f := fromFile
				r := fromRank

				if diffFile > 0 && diffRank > 0 {
					f++
					r++
					for range diffFile - 1 {
						if Board[r][f].PieceType != board.NoPieceType {
							return false
						}
						f++
						r++
					}
				} else if diffFile > 0 && diffRank < 0 {
					f++
					r--
					for range diffFile - 1 {
						if Board[r][f].PieceType != board.NoPieceType {
							return false
						}
						f++
						r--
					}
				} else if diffFile < 0 && diffRank < 0 {
					f--
					r--
					for range diffFile*(-1) - 1 {
						if Board[r][f].PieceType != board.NoPieceType {
							return false
						}
						f--
						r--
					}
				} else if diffFile < 0 && diffRank > 0 {
					f--
					r++
					for range diffRank - 1 {
						if Board[r][f].PieceType != board.NoPieceType {
							return false
						}
						f--
						r++
					}
				}
				return true
			}

		}

	case board.King:
		if (math.Abs(float64(toFile-fromFile)) <= 1 && math.Abs(float64(toRank-fromRank)) <= 1) &&
			(Board[toRank][toFile].PieceColor != pieceColor) {
			return true
		}
	case board.NoPieceType:
		return false
	}
	return false
}

