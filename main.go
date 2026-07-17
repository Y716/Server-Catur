package main

import (
	"fmt"
)

type PieceType int

const (
	NoPieceType PieceType = iota
	King
	Queen
	Rook
	Bishop
	Knight
	Pawn
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
	Board := [8][8]Piece{}
	return Board
}

func main() {
	Board := NewBoard()
	fmt.Println(Board)
}
