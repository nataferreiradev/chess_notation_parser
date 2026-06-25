// Package chess interpreta a notação algébrica de lances de xadrez.
//
// A função principal é Parse, que recebe uma jogada como "Bxe4+" e devolve um
// Move estruturado. Move.String reconstrói a notação correspondente.
package chess

import (
	"errors"
	"fmt"
)

var (
	errInvalidFile = errors.New("valor invalido para file")
	errInvalidRank = errors.New("valor invalido para rank")
)

// Parse interpreta uma jogada em notação algébrica e devolve o Move
// correspondente. Retorna erro se a notação for inválida.
func Parse(cmd string) (Move, error) {
	if len(cmd) < 2 {
		return Move{}, fmt.Errorf("notação muito curta: %q", cmd)
	}

	check := false
	checkmate := false

	switch cmd[len(cmd)-1] {
	case '+':
		check = true
		cmd = cmd[:len(cmd)-1]
	case '#':
		checkmate = true
		cmd = cmd[:len(cmd)-1]
	}

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

	return Move{
		Piece:     piece,
		From:      from,
		Capture:   capture,
		Check:     check,
		Checkmate: checkmate,
		To: Square{
			File: file,
			Rank: rank,
		},
	}, nil
}
