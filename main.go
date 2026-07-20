package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func boardRepToCompRep(number int) int {
	// Change board representation to array representation: y = -x + 8
	return number*(-1) + 8
}

func MovePiece(from string, to string, Board [8][8]Piece) {
	fromFile := alphaToNum(string(from[0])) - 1
	fromRank, _ := strconv.Atoi(string(from[1]))
	toFile := alphaToNum(string(to[0])) - 1
	toRank, _ := strconv.Atoi(string(to[1]))

	movedPiece := Board[boardRepToCompRep(fromRank)][fromFile]
	Board[boardRepToCompRep(fromRank)][fromFile] = Piece{NoPieceType, NoPieceColor}
	Board[boardRepToCompRep(toRank)][toFile] = movedPiece

	PrintBoard(Board)
}

func main() {
	Board := NewBoard()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Your move: ")
	if scanner.Scan() {
		input := scanner.Text()
		notations := strings.Split(input, " ")
		MovePiece(notations[0], notations[1], Board)
		PrintBoard(Board)
	}
}
