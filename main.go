package main

import (
	"fmt"
)

type PieceType int

const (
	NoPieceType PieceType = iota
	Pawn
	Bishop
	Knight
	Rook
	Queen
	King
)

type PieceColor int

const (
	NoPieceColor PieceColor = iota
	Black
	White
)

type Piece struct {
	pieceType  PieceType
	pieceColor PieceColor
}

func NewBoard() [8][8]Piece {
	Board := [8][8]Piece{
		{{Rook, Black}, {Knight, Black}, {Bishop, Black}, {Queen, Black}, {King, Black}, {Bishop, Black}, {Knight, Black}, {Rook, Black}},
		{{Pawn, Black}, {Pawn, Black}, {Pawn, Black}, {Pawn, Black}, {Pawn, Black}, {Pawn, Black}, {Pawn, Black}, {Pawn, Black}},
		{{NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}},
		{{NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}},
		{{NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}},
		{{NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}, {NoPieceType, NoPieceColor}},
		{{Pawn, White}, {Pawn, White}, {Pawn, White}, {Pawn, White}, {Pawn, White}, {Pawn, White}, {Pawn, White}, {Pawn, White}},
		{{Rook, White}, {Knight, White}, {Bishop, White}, {Queen, White}, {King, White}, {Bishop, White}, {Knight, White}, {Rook, White}},
	}
	return Board
}

func getUniCodePiece(piece Piece) string {
	pieceMap := map[Piece]string{
		{NoPieceType, NoPieceColor}: "\u25A1",
		{Pawn, Black}:               "\u265F",
		{Bishop, Black}:             "\u265D",
		{Knight, Black}:             "\u265E",
		{Rook, Black}:               "\u265C",
		{Queen, Black}:              "\u265B",
		{King, Black}:               "\u265A",
		{Pawn, White}:               "\u2659",
		{Bishop, White}:             "\u2657",
		{Knight, White}:             "\u2658",
		{Rook, White}:               "\u2656",
		{Queen, White}:              "\u2655",
		{King, White}:               "\u2654",
	}

	return pieceMap[piece]
}

func PrintBoard(Board [8][8]Piece) {
	for i := 0; i < len(Board); i++ {
		for j := 0; j < len(Board); j++ {
			uniCodePiece := getUniCodePiece(Board[i][j])
			fmt.Print(uniCodePiece + " ")
		}
		fmt.Println()
	}
}

func main() {
	Board := NewBoard()
	PrintBoard(Board)
}
