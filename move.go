package chess

import "strings"

// CastleSide indica o lado do roque, quando o lance é um roque.
type CastleSide byte

const (
	// CastleNone é o valor de um lance que não é roque.
	CastleNone CastleSide = iota
	// CastleKingside é o roque pequeno (O-O).
	CastleKingside
	// CastleQueenside é o roque grande (O-O-O).
	CastleQueenside
)

// String devolve a notação do roque, ou string vazia se não houver roque.
func (c CastleSide) String() string {
	switch c {
	case CastleKingside:
		return "O-O"
	case CastleQueenside:
		return "O-O-O"
	}
	return ""
}

// Move representa um lance já interpretado a partir da notação algébrica.
//
// From costuma vir parcialmente preenchido: a notação algébrica só informa a
// casa de origem quando ela é necessária para desambiguar. Um componente não
// informado vale -1, e cabe a quem tem o tabuleiro resolver a origem real.
//
// Num roque, Castle diz o lado, Piece é King e From e To vêm ambos {-1, -1}:
// as casas dependem da cor de quem joga, que a notação não carrega.
type Move struct {
	Piece     Piece
	From      Square
	To        Square
	Capture   bool
	Check     bool
	Checkmate bool
	Castle    CastleSide
	// Promotion é a peça escolhida na promoção do peão, ou 0 se o lance não
	// for uma promoção.
	Promotion Piece
}

// String reconstrói a notação algébrica do lance.
//
// As anotações de comentário (!, ?) não são guardadas em Move, então não
// reaparecem aqui: Parse("e4!?").String() é "e4".
func (m Move) String() string {
	var sb strings.Builder

	if m.Castle != CastleNone {
		sb.WriteString(m.Castle.String())
		writeCheckSuffix(&sb, m)
		return sb.String()
	}

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

	if m.Promotion != 0 {
		sb.WriteByte('=')
		sb.WriteByte(byte(m.Promotion))
	}

	writeCheckSuffix(&sb, m)

	return sb.String()
}

func writeCheckSuffix(sb *strings.Builder, m Move) {
	if m.Check {
		sb.WriteByte('+')
	}
	if m.Checkmate {
		sb.WriteByte('#')
	}
}
