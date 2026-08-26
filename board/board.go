package board

import (
	"fmt"
	"strings"
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

type File int

const (
	a File = iota
	b
	c
	d
	e
	f
	g
	h
)

type Piece struct {
	PieceType  PieceType
	PieceColor PieceColor
}

func NewBoard() *[8][8]Piece {
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
	return &Board
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
func PrintBoard(Board [8][8]Piece) string { 
	var boardState strings.Builder
	boardState.WriteString("  ")
	for ch := 'A'; ch <= 'H'; ch++{
		fmt.Fprintf(&boardState, "%c ", ch)
	}
	boardState.WriteString("\n")
	for i := 0; i < len(Board); i++ {
		for j := 0; j < len(Board); j++ {
			if j == 0{
				fmt.Fprintf(&boardState, "%d ", (i*-1)+8)
			}
			uniCodePiece := getUniCodePiece(Board[i][j])
			fmt.Fprintf(&boardState, "%s ", uniCodePiece)
			if j == 7{
				fmt.Fprintf(&boardState, "%d ", (i*-1)+8)
			}
		}
		
		fmt.Fprintln(&boardState, "")

	}
	fmt.Fprint(&boardState, "  ")
	for ch := 'A'; ch <= 'H'; ch++{
		fmt.Fprintf(&boardState, "%c ", ch)
	}
	fmt.Fprintln(&boardState, "")

	fmt.Printf("%s", &boardState)
	return boardState.String()
}

// func PrintBoard(Board [8][8]Piece) string {
// 	for ch := 'A'; ch <= 'H'; ch++{
// 		fmt.Printf("%c ", ch)
// 	}
// 	fmt.Println()
// 	for i := 0; i < len(Board); i++ {
// 		for j := 0; j < len(Board); j++ {
// 			if j == 0{
// 				fmt.Printf("%d ", (i*-1)+8)
// 			}
// 			uniCodePiece := getUniCodePiece(Board[i][j])
// 			fmt.Print(uniCodePiece + " ")
// 			if j == 7{
// 				fmt.Printf("%d", (i*-1)+8)
// 			}
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Print("  ")
// 	for ch := 'A'; ch <= 'H'; ch++{
// 		fmt.Printf("%c ", ch)
// 	}
// 	fmt.Println()
//
// }
