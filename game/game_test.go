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

	MovePiece(testBoard, fromSquare, toSquare, true)

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

func TestColorTurn(t *testing.T) {
	// Test MovePiece func
	tests := map[string]struct {
		colorFlag bool
		fromSquare string
		toSquare   string
		result     board.Piece
	}{
		"BlackMoveinWhiteTurn": {
			colorFlag: true,
			fromSquare: "e7",
			toSquare:   "e6",
			result: board.Piece{
				PieceType:  board.NoPieceType,
				PieceColor: board.NoPieceColor,
			},
		},
		"WhiteMoveinBlackTurn": {
			colorFlag: false,
			fromSquare: "e2",
			toSquare:   "e3",
			result: board.Piece{
				PieceType:  board.NoPieceType,
				PieceColor: board.NoPieceColor,
			},
		},
		"WhiteMoveinWhiteTurn": {
			colorFlag: true,
			fromSquare: "e2",
			toSquare:   "e3",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
		"BlackMoveinBlackTurn": {
			colorFlag: false,
			fromSquare: "e7",
			toSquare:   "e6",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.Black,
			},
		},
	}

	// fromFile, fromRank := boardRepToCompRep(fromSquare)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testBoard := board.NewBoard()
			MovePiece(testBoard, test.fromSquare, test.toSquare, test.colorFlag)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
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
			MovePiece(testBoard, test.fromSquare, test.toSquare, true)
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

			MovePiece(&testBoard, test.fromSquare, test.toSquare, true)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
	}
}

func TestMoveKnight(t *testing.T) {
	// Test MovePiece func
	tests := map[string]struct {
		fromSquare string
		toSquare   string
		result     board.Piece
	}{
		"E4toC5": {
			fromSquare: "e4",
			toSquare:   "c5",
			result: board.Piece{
				PieceType:  board.Knight,
				PieceColor: board.White,
			},
		},
		"E4toG5Take": {
			fromSquare: "e4",
			toSquare:   "g5",
			result: board.Piece{
				PieceType:  board.Knight,
				PieceColor: board.White,
			},
		},
		"E4toG3Blocked": {
			fromSquare: "e4",
			toSquare:   "g3",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
	}

	// fromFile, fromRank := boardRepToCompRep(fromSquare)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testBoard := [8][8]board.Piece{}
			testBoard[4][4] = board.Piece{PieceType: board.Knight, PieceColor: board.White} //Knight Putih di E4
			testBoard[3][6] = board.Piece{PieceType: board.Pawn, PieceColor: board.Black}   //Pawn Hitam di G5
			testBoard[5][6] = board.Piece{PieceType: board.Pawn, PieceColor: board.White}   //Pawn Putih di G3

			MovePiece(&testBoard, test.fromSquare, test.toSquare, true)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
	}
}

func TestMoveBishop(t *testing.T) {
	// Test MovePiece func
	tests := map[string]struct {
		fromSquare string
		toSquare   string
		result     board.Piece
	}{
		"E4toH7": {
			fromSquare: "e4",
			toSquare:   "h7",
			result: board.Piece{
				PieceType:  board.Bishop,
				PieceColor: board.White,
			},
		},
		"E4toA8Take": {
			fromSquare: "e4",
			toSquare:   "a8",
			result: board.Piece{
				PieceType:  board.Bishop,
				PieceColor: board.White,
			},
		},
		"E4toB1": {
			fromSquare: "e4",
			toSquare:   "b1",
			result: board.Piece{
				PieceType:  board.Bishop,
				PieceColor: board.White,
			},
		},
		"E4toG2": {
			fromSquare: "e4",
			toSquare:   "g2",
			result: board.Piece{
				PieceType:  board.Bishop,
				PieceColor: board.White,
			},
		},
		"E4toH1Blocked": {
			fromSquare: "e4",
			toSquare:   "h1",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
	}

	// fromFile, fromRank := boardRepToCompRep(fromSquare)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testBoard := [8][8]board.Piece{}
			testBoard[4][4] = board.Piece{PieceType: board.Bishop, PieceColor: board.White} //Bishop Putih di E4
			testBoard[0][0] = board.Piece{PieceType: board.Pawn, PieceColor: board.Black}   //Pawn Hitam di A8
			testBoard[7][7] = board.Piece{PieceType: board.Pawn, PieceColor: board.White}   //Pawn Putih di H1

			MovePiece(&testBoard, test.fromSquare, test.toSquare, true)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
	}
}

func TestMoveQueen(t *testing.T) {
	// Test MovePiece func
	tests := map[string]struct {
		fromSquare string
		toSquare   string
		result     board.Piece
	}{
		"E4toH7": {
			fromSquare: "e4",
			toSquare:   "h7",
			result: board.Piece{
				PieceType:  board.Queen,
				PieceColor: board.White,
			},
		},
		"E4toA8Take": {
			fromSquare: "e4",
			toSquare:   "a8",
			result: board.Piece{
				PieceType:  board.Queen,
				PieceColor: board.White,
			},
		},
		"E4toB1": {
			fromSquare: "e4",
			toSquare:   "b1",
			result: board.Piece{
				PieceType:  board.Queen,
				PieceColor: board.White,
			},
		},
		"E4toG2": {
			fromSquare: "e4",
			toSquare:   "g2",
			result: board.Piece{
				PieceType:  board.Queen,
				PieceColor: board.White,
			},
		},
		"E4toH1Blocked": {
			fromSquare: "e4",
			toSquare:   "h1",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
		"E4toE1": {
			fromSquare: "e4",
			toSquare:   "e1",
			result: board.Piece{
				PieceType:  board.Queen,
				PieceColor: board.White,
			},
		},
		"E4toB5Take": {
			fromSquare: "e4",
			toSquare:   "b4",
			result: board.Piece{
				PieceType:  board.Queen,
				PieceColor: board.White,
			},
		},
		"E4toG5Blocked": {
			fromSquare: "e4",
			toSquare:   "g4",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
		"E4toA5Blocked": {
			fromSquare: "e4",
			toSquare:   "a4",
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
			testBoard[4][4] = board.Piece{PieceType: board.Queen, PieceColor: board.White} //Queen Putih di E4
			testBoard[0][0] = board.Piece{PieceType: board.Pawn, PieceColor: board.Black}  //Pawn Hitam di A8
			testBoard[7][7] = board.Piece{PieceType: board.Pawn, PieceColor: board.White}  //Pawn Putih di H1

			testBoard[4][1] = board.Piece{PieceType: board.Pawn, PieceColor: board.Black} //Pawn Hitam di B4
			testBoard[4][6] = board.Piece{PieceType: board.Pawn, PieceColor: board.White} //Pawn Putih di G4

			MovePiece(&testBoard, test.fromSquare, test.toSquare, true)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
	}
}

func TestMoveKing(t *testing.T) {
	// Test MovePiece func
	tests := map[string]struct {
		fromSquare string
		toSquare   string
		result     board.Piece
	}{
		"E4toF4BlockedByCheck": {
			fromSquare: "e4",
			toSquare:   "f4",
			result: board.Piece{
				PieceType:  board.NoPieceType,
				PieceColor: board.NoPieceColor,
			},
		},
		"E4toF5": {
			fromSquare: "e4",
			toSquare:   "f5",
			result: board.Piece{
				PieceType:  board.King,
				PieceColor: board.White,
			},
		},
		"E4toE5Take": {
			fromSquare: "e4",
			toSquare:   "e5",
			result: board.Piece{
				PieceType:  board.King,
				PieceColor: board.White,
			},
		},
		"E4toE3Blocked": {
			fromSquare: "e4",
			toSquare:   "e3",
			result: board.Piece{
				PieceType:  board.Pawn,
				PieceColor: board.White,
			},
		},
	}

	// fromFile, fromRank := boardRepToCompRep(fromSquare)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testBoard := [8][8]board.Piece{}
			testBoard[4][4] = board.Piece{PieceType: board.King, PieceColor: board.White} //King Putih di E4
			testBoard[3][4] = board.Piece{PieceType: board.Pawn, PieceColor: board.Black} //Pawn Hitam di E5
			testBoard[5][4] = board.Piece{PieceType: board.Pawn, PieceColor: board.White} //Pawn Putih di E3

			MovePiece(&testBoard, test.fromSquare, test.toSquare, true)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
	}
}

func TestKingInCheck(t *testing.T){
	// Test MovePiece func
	tests := map[string]struct {
		fromSquare string
		toSquare   string
		result     board.Piece
	}{
		"E4toE5Blocked": {
			fromSquare: "e4",
			toSquare:   "e5",
			result: board.Piece{
				PieceType:  board.NoPieceType,
				PieceColor: board.NoPieceColor,
			},
		},
		"E4toF4": {
			fromSquare: "e4",
			toSquare:   "f4",
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
			testBoard := [8][8]board.Piece{}
			testBoard[4][4] = board.Piece{PieceType: board.King, PieceColor: board.White} //King Putih di E4
			testBoard[3][7] = board.Piece{PieceType: board.Rook, PieceColor: board.Black} //Pawn Hitam di H5
			testBoard[5][7] = board.Piece{PieceType: board.Rook, PieceColor: board.Black} //Pawn Putih di H3

			MovePiece(&testBoard, test.fromSquare, test.toSquare, true)
			toFile, toRank := boardRepToCompRep(test.toSquare)
			gotPiece := testBoard[toRank][toFile]
			if expeceted := test.result; gotPiece != expeceted {
				t.Fatalf("returned %+v expeceted %+v", gotPiece, expeceted)
			}
		})
	}
}
