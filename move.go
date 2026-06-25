package chess

import "strings"

// Move representa um lance já interpretado a partir da notação algébrica.
type Move struct {
	Piece     Piece
	From      Square
	To        Square
	Capture   bool
	Check     bool
	Checkmate bool
}

// String reconstrói a notação algébrica do lance.
func (m Move) String() string {
	var sb strings.Builder

	if m.Piece != Pawn {
		sb.WriteByte(byte(m.Piece))
	}

	if m.From.File >= 0 {
		sb.WriteByte(byte('a' + m.From.File))
	}
	if m.From.Rank >= 0 {
		sb.WriteByte(byte('1' + m.From.Rank))
	}

	if m.Capture {
		sb.WriteByte('x')
	}

	sb.WriteByte(byte('a' + m.To.File))
	sb.WriteByte(byte('1' + m.To.Rank))

	if m.Check {
		sb.WriteByte('+')
	}

	if m.Checkmate {
		sb.WriteByte('#')
	}

	return sb.String()
}
