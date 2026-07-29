package game

import (
	"github.com/Y716/Server-Catur/board"
)

func isValidMove(Board *[8][8]board.Piece, from string, to string, colorFlag bool) bool {
	fromFile, fromRank := boardRepToCompRep(from)
	toFile, toRank := boardRepToCompRep(to)
	pieceType := Board[fromRank][fromFile].PieceType
	pieceColor := Board[fromRank][fromFile].PieceColor

	if colorFlag && pieceColor != board.White || !colorFlag && pieceColor != board.Black {
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
				Board[toRank][toFile].PieceColor != pieceColor {
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
		// Rook hanya bisa gerak horizontal dan vertikal (File beda tapi Rank sama atau Rank beda File sama (antara fromSquare dan toSquare))
		// Rook terhalang oleh bidak ke arah toSquarenya.Jika ada bidak, maka invalid.
		// Rook dapat memakan bidak yang berlawanan dengan warnanya
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

			if Board[toRank][toFile].PieceColor == pieceColor {
				break
			}
			return true

		}
	}

	return false
}
