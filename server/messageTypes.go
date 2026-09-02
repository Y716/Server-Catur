package server

import (
	"github.com/Y716/Server-Catur/board"
)

type Message struct{
	Type string `json:"type"`
}

type MoveMessage struct {
	Type string `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`
}

type WarningMessage struct{
	Type string `json:"type"`
	Message string `json:"message"`
}
type BoardMessage struct {
	Type  string            `json:"type"`
	Board [8][8]board.Piece `json:"board"`
}
type TurnMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ColorMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
