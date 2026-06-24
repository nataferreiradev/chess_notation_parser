package main

import (
	"fmt"
	"strings"
)

func main() {
	move, err := parse("Bxe4")
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Println(move.toString())
	fmt.Printf("%c\n", move.piece)
	fmt.Println(move.from.File)
	fmt.Println(move.from.Rank)
	fmt.Println(move.capture)
	fmt.Println(move.to.File)
	fmt.Println(move.to.Rank)
}

type Piece byte

const (
	Pawn   Piece = 'P'
	Knight Piece = 'N'
	Bishop Piece = 'B'
	Rook   Piece = 'R'
	Queen  Piece = 'Q'
	King   Piece = 'K'
)

type Square struct {
	File int // coluna: a=0 .. h=7  (-1 = não especificado)
	Rank int // linha:  1=0 .. 8=7  (-1 = não especificado)
}

type Move struct {
	piece   Piece
	from    Square
	to      Square
	capture bool
}

func (m Move) toString() string {
	var sb strings.Builder

	if m.piece != Pawn {
		sb.WriteByte(byte(m.piece))
	}

	if m.from.File >= 0 {
		sb.WriteByte(byte('a' + m.from.File))
	}
	if m.from.Rank >= 0 {
		sb.WriteByte(byte('1' + m.from.Rank))
	}

	if m.capture {
		sb.WriteByte('x')
	}

	sb.WriteByte(byte('a' + m.to.File))
	sb.WriteByte(byte('1' + m.to.Rank))

	return sb.String()
}

func calculateFileValue(charByte byte) (int, error) {
	value := int(charByte - 'a')
	if value < 0 || value > 7 {
		return -1, fmt.Errorf("valor invalido para file")
	}
	return value, nil
}

func calculateRankValue(charByte byte) (int, error) {
	value := int(charByte - '1')
	if value < 0 || value > 7 {
		return -1, fmt.Errorf("valor invalido para rank")
	}
	return value, nil
}

func decidePiece(charByte byte) Piece {
	switch Piece(charByte) {
	case King, Queen, Bishop, Rook, Knight:
		return Piece(charByte)
	default:
		return Pawn
	}
}

func parse(cmd string) (Move, error) {
	if len(cmd) < 2 {
		return Move{}, fmt.Errorf("notação muito curta: %q", cmd)
	}

	dest := cmd[len(cmd)-2:]
	prefix := cmd[:len(cmd)-2]

	file, err := calculateFileValue(dest[0])
	if err != nil {
		return Move{}, err
	}

	rank, err := calculateRankValue(dest[1])
	if err != nil {
		return Move{}, err
	}

	from := Square{File: -1, Rank: -1}
	capture := false

	if len(prefix) > 0 && prefix[len(prefix)-1] == 'x' {
		capture = true
		prefix = prefix[:len(prefix)-1]
	}

	piece := Pawn
	if len(prefix) > 0 {
		p := decidePiece(prefix[0])
		if p != Pawn {
			piece = p
		} else {
			f, err := calculateFileValue(prefix[0])
			if err != nil {
				return Move{}, err
			}
			from.File = f
		}
	}

	isPawnMove := piece == Pawn

	if !isPawnMove && len(cmd) < 3 {
		return Move{}, fmt.Errorf("notação muito curta: %q", cmd)
	}

	return Move{
		piece:   piece,
		from:    from,
		capture: capture,
		to: Square{
			File: file,
			Rank: rank,
		},
	}, nil

}
