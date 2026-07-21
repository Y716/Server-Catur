package game

import (
	"testing"

	"github.com/Y716/Server-Catur/board"
)

func TestGame(t *testing.T) {
	// Test MovePiece func
	testBoard := board.NewBoard()
	fromSquare := "e2"
	toSquare := "e4"

	MovePiece(fromSquare, toSquare, &testBoard)

	fromFile, fromRank := boardRepToCompRep(fromSquare)
	toFile, toRank := boardRepToCompRep(toSquare)

	emptySquare := board.Piece{
		PieceType:  board.NoPieceType,
		PieceColor: board.NoPieceColor,
	}
	if testBoard[fromRank][fromFile] != emptySquare {
		t.Fatalf("returned %+v expeceted %+v", testBoard[fromRank][fromFile], emptySquare)
	}

	whitePawn := board.Piece{
		PieceType:  board.Pawn,
		PieceColor: board.White,
	}

	if testBoard[toRank][toFile] != whitePawn {
		t.Fatalf("returned %+v expeceted %+v", testBoard[toRank][toFile], whitePawn)
	}
}
