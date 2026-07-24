package game

import (
	"github.com/Y716/Server-Catur/board"
)

func isValidMove(Board *[8][8]board.Piece, from string, to string) bool {
	fromFile, fromRank := boardRepToCompRep(from)
	toFile, toRank := boardRepToCompRep(to)
	pieceType := Board[fromRank][fromFile].PieceType
	PieceColor := Board[fromRank][fromFile].PieceColor

	switch pieceType {
	case board.Pawn:
		if Board[toRank][toFile].PieceType == board.NoPieceType {
			if fromFile == toFile {
				if fromRank == 1 && PieceColor == board.Black || fromRank == 6 && PieceColor == board.White {
					diff := fromRank - toRank
					if ((diff == -1 || diff == -2) && PieceColor == board.Black) || ((diff == 1 || diff == 2) && PieceColor == board.White) {

						if (diff == 2 && Board[fromRank-1][fromFile].PieceType != board.NoPieceType) || (diff == -2 && Board[fromRank+1][fromFile].PieceType != board.NoPieceType) {
							break
						}
						return true
					} else {
						break
					}
				} else {
					diff := fromRank - toRank
					if diff == -1 && PieceColor == board.Black || diff == 1 && PieceColor == board.White {
						return true
					} else {
						break
					}
				}
			} else {
				break
			}
		} else {
			break
		}
	}

	return false
}
