package game

import (
	"testing"

	"github.com/Y716/Server-Catur/board"
)

func TestMovePiece(t *testing.T) {
	// Test MovePiece func
	testBoard := board.NewBoard()
	fromSquare := "e2"
	toSquare := "e4"

	MovePiece(&testBoard, fromSquare, toSquare)

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

func TestMovePawn(t *testing.T) {
	// Test MovePiece func
	tests := map[string]struct {
		fromSquare string
		toSquare   string
		result     board.Piece
	}{
		"E2toE3": {
			fromSquare: "e2",
			toSquare:   "e3",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
		"E2toE4": {
			fromSquare: "e2",
			toSquare:   "e4",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
		"E2toE5": {
			fromSquare: "e2",
			toSquare:   "e5",
			result: board.Piece{
				PieceType:  board.NoPieceType,
				PieceColor: board.NoPieceColor,
			},
		},
		"E2toE1": {
			fromSquare: "e2",
			toSquare:   "e1",
			result: board.Piece{
				PieceType:  board.King,
				PieceColor: board.White,
			},
		},
	}

	// fromFile, fromRank := boardRepToCompRep(fromSquare)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testBoard := board.NewBoard()
			MovePiece(&testBoard, test.fromSquare, test.toSquare)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
	}
}

func TestMoveRook(t *testing.T) {
	// Test MovePiece func
	tests := map[string]struct {
		fromSquare string
		toSquare   string
		result     board.Piece
	}{
		"E5toE1": {
			fromSquare: "e5",
			toSquare:   "e1",
			result: board.Piece{
				PieceType:  board.Rook,
				PieceColor: board.White,
			},
		},
		"E5toB5Take": {
			fromSquare: "e5",
			toSquare:   "b5",
			result: board.Piece{
				PieceType:  board.Rook,
				PieceColor: board.White,
			},
		},
		"E5toG5Blocked": {
			fromSquare: "e5",
			toSquare:   "g5",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
		"E5toA5Blocked": {
			fromSquare: "e5",
			toSquare:   "a5",
			result: board.Piece{
				PieceType:  board.NoPieceType,
				PieceColor: board.NoPieceColor,
			},
		},
	}

	// fromFile, fromRank := boardRepToCompRep(fromSquare)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testBoard := [8][8]board.Piece{}
			testBoard[3][4] = board.Piece{PieceType: board.Rook, PieceColor: board.White} //Rook Putih di E5
			testBoard[3][1] = board.Piece{PieceType: board.Pawn, PieceColor: board.Black} //Pawn Hitam di B5
			testBoard[3][6] = board.Piece{PieceType: board.Pawn, PieceColor: board.White} //Pawn Putih di G5

			MovePiece(&testBoard, test.fromSquare, test.toSquare)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
	}
}
