package game

import (
	"strconv"
	"strings"

	"github.com/Y716/Server-Catur/board"
)

func alphaToNum(s string) int {
	// Standardize to uppercase to easily handle mixed input
	s = strings.ToUpper(s)

	result := 0
	for i := 0; i < len(s); i++ {
		// Convert character to 1-26 value
		charValue := int(s[i] - 'A' + 1)

		// Handle multi-letter lists (like Excel columns: A=1, Z=26, AA=27)
		result = result*26 + charValue
	}
	return result
}

func boardRepToCompRep(square string) (int, int) {
	// Change board representation to array representation: y = -x + 8
	file := alphaToNum(string(square[0])) - 1
	rank, _ := strconv.Atoi(string(square[1]))

	return file, rank*(-1) + 8
}

func MovePiece(from string, to string, Board *[8][8]board.Piece) {
	fromFile, fromRank := boardRepToCompRep(from)
	toFile, toRank := boardRepToCompRep(to)

	movedPiece := Board[fromRank][fromFile]
	Board[fromRank][fromFile] = board.Piece{
		PieceType:  board.NoPieceType,
		PieceColor: board.NoPieceColor,
	}
	Board[toRank][toFile] = movedPiece

	board.PrintBoard(*Board)
}
