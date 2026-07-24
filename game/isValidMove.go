package game

import (
	"github.com/Y716/Server-Catur/board"
)

func isValidMove(Board *[8][8]board.Piece, from string, to string) bool {
	fromFile, fromRank := boardRepToCompRep(from)
	toFile, toRank := boardRepToCompRep(to)
	pieceType := Board[fromRank][fromFile].PieceType
	pieceColor := Board[fromRank][fromFile].PieceColor

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
	}

	return false
}
