package board

import "testing"

func TestBoard(t *testing.T) {
	//Test Newboard dan PrintBoard func
	testBoard := NewBoard()

	tests := map[string]struct {
		input  Piece
		result Piece
	}{
		"White King": {
			input: testBoard[7][4],
			result: Piece{
				PieceType:  King,
				PieceColor: White},
		},
		"Black King": {
			input: testBoard[0][4],
			result: Piece{
				PieceType:  King,
				PieceColor: Black},
		},
		"White Queen": {
			input: testBoard[7][3],
			result: Piece{
				PieceType:  Queen,
				PieceColor: White},
		},
		"Black Queen": {
			input: testBoard[0][3],
			result: Piece{
				PieceType:  Queen,
				PieceColor: Black},
		},
		"White King Rook": {
			input: testBoard[7][7],
			result: Piece{
				PieceType:  Rook,
				PieceColor: White},
		},
		"Black King Rook": {
			input: testBoard[0][7],
			result: Piece{
				PieceType:  Rook,
				PieceColor: Black},
		},
		"White Queen Rook": {
			input: testBoard[7][0],
			result: Piece{
				PieceType:  Rook,
				PieceColor: White},
		},
		"Black Queen Rook": {
			input: testBoard[0][0],
			result: Piece{
				PieceType:  Rook,
				PieceColor: Black},
		},
		"White King Knight": {
			input: testBoard[7][6],
			result: Piece{
				PieceType:  Knight,
				PieceColor: White},
		},
		"Black King Knight": {
			input: testBoard[0][6],
			result: Piece{
				PieceType:  Knight,
				PieceColor: Black},
		},
		"White Queen Knight": {
			input: testBoard[7][1],
			result: Piece{
				PieceType:  Knight,
				PieceColor: White},
		},
		"Black Queen Knight": {
			input: testBoard[0][1],
			result: Piece{
				PieceType:  Knight,
				PieceColor: Black},
		},
		"White King Bishop": {
			input: testBoard[7][5],
			result: Piece{
				PieceType:  Bishop,
				PieceColor: White},
		},
		"Black King Bishop": {
			input: testBoard[0][5],
			result: Piece{
				PieceType:  Bishop,
				PieceColor: Black},
		},
		"White Queen Bishop": {
			input: testBoard[7][2],
			result: Piece{
				PieceType:  Bishop,
				PieceColor: White},
		},
		"Black Queen Bishop": {
			input: testBoard[0][2],
			result: Piece{
				PieceType:  Bishop,
				PieceColor: Black},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expeceted := test.input, test.result; got != expeceted {
				t.Fatalf("returned %+v expeceted %+v", got, expeceted)
			}
		})
	}

	//Check Pawns
	blackPawn := Piece{
		PieceType:  Pawn,
		PieceColor: Black,
	}
	for i := 0; i < 8; i++ {
		if testBoard[1][i] != blackPawn {
			t.Fatalf("returned %+v expeceted %+v at testBoard[1][%d]", testBoard[1][i], blackPawn, i)
		}
	}

	whitePawn := Piece{
		PieceType:  Pawn,
		PieceColor: White,
	}
	for i := 0; i < 8; i++ {
		if testBoard[6][i] != whitePawn {
			t.Fatalf("returned %+v expeceted %+v at testBoard[6][%d]", testBoard[1][i], whitePawn, i)
		}
	}

	//Check Empty Squares
	emptySquare := Piece{
		PieceType:  NoPieceType,
		PieceColor: NoPieceColor,
	}
	for i := 0; i < 8; i++ {
		for j := 2; j < 6; j++ {
			if testBoard[j][i] != emptySquare {
				t.Fatalf("returned %+v expeceted %+v at testBoard[1][%d]", testBoard[1][i], emptySquare, i)
			}
		}

	}

}
