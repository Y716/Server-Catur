package game

import (
	"fmt"
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
func MovePiece(Board *[8][8]board.Piece, from string, to string, colorFlag bool) bool {
	fromFile, fromRank := boardRepToCompRep(from)
	toFile, toRank := boardRepToCompRep(to)
	
	pieceColor := Board[fromRank][fromFile].PieceColor

	if colorFlag && pieceColor != board.White || !colorFlag && pieceColor != board.Black {
		return false
	}

	if !isValidMove(Board, fromFile, toFile, fromRank, toRank) {
		fmt.Println("Invalid Move")
		return false
	}

	simulationBoard := *Board

	if IsKingInCheckAfterMove(simulationBoard, fromRank, fromFile, toRank, toFile, colorFlag){
		fmt.Println("King is in Check!")
		return false
	}
	movedPiece := simulationBoard[fromRank][fromFile]

	Board[fromRank][fromFile] = board.Piece{
		PieceType:  board.NoPieceType,
		PieceColor: board.NoPieceColor,
	}
	Board[toRank][toFile] = movedPiece
	return true
}
