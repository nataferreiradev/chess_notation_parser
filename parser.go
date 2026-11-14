// Package chess interpreta a notação algébrica de lances de xadrez.
//
// A função principal é Parse, que recebe uma jogada como "Bxe4+" e devolve um
// Move estruturado. Move.String reconstrói a notação correspondente.
//
// O parser é puramente sintático: não conhece o tabuleiro, então não sabe se a
// peça existe, se alcança o destino ou se o lance é legal. Ele garante apenas
// que a notação está bem formada e entrega os componentes que conseguiu ler.
package chess

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errInvalidFile = errors.New("valor invalido para file")
	errInvalidRank = errors.New("valor invalido para rank")
)

// Parse interpreta uma jogada em notação algébrica e devolve o Move
// correspondente. Retorna erro se a notação for inválida.
//
// Aceita as formas usuais: lance simples ("e4", "Nf3"), captura ("exd5",
// "Bxe4"), desambiguação por coluna, por linha ou por casa inteira ("Nbd2",
// "R1a3", "Qh4e1"), roque ("O-O", "O-O-O"), promoção ("e8=Q", "exd8=N+") e os
// sufixos de xeque e mate ("+", "#"). Anotações de comentário ("!", "?") são
// toleradas e descartadas.
func Parse(cmd string) (Move, error) {
	cmd = strings.TrimSpace(cmd)

	// Anotações de comentário não mudam o lance; descartamos antes de tudo.
	cmd = strings.TrimRight(cmd, "!?")

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

	if castle := parseCastle(cmd); castle != CastleNone {
		return Move{
			Piece:     King,
			Castle:    castle,
			From:      Square{File: -1, Rank: -1},
			To:        Square{File: -1, Rank: -1},
			Check:     check,
			Checkmate: checkmate,
		}, nil
	}

	promotion := Piece(0)
	if idx := strings.IndexByte(cmd, '='); idx >= 0 {
		if idx != len(cmd)-2 {
			return Move{}, fmt.Errorf("promoção malformada: %q", cmd)
		}
		p := decidePiece(upper(cmd[idx+1]))
		switch p {
		case Queen, Rook, Bishop, Knight:
			promotion = p
		default:
			return Move{}, fmt.Errorf("não se pode promover a %q", string(cmd[idx+1]))
		}
		cmd = cmd[:idx]
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

	capture := false
	if len(prefix) > 0 && prefix[len(prefix)-1] == 'x' {
		capture = true
		prefix = prefix[:len(prefix)-1]
	}

	// A letra da peça, quando existe, é sempre maiúscula e vem primeiro; as
	// colunas são minúsculas, então não há como confundir "B" (bispo) com "b"
	// (coluna b).
	piece := Pawn
	if len(prefix) > 0 {
		if p := decidePiece(prefix[0]); p != Pawn {
			piece = p
			prefix = prefix[1:]
		}
	}

	// O que sobrou é a desambiguação da origem: "", coluna, linha ou as duas.
	from := Square{File: -1, Rank: -1}
	if len(prefix) > 0 {
		if f, err := calculateFileValue(prefix[0]); err == nil {
			from.File = f
			prefix = prefix[1:]
		}
	}
	if len(prefix) > 0 {
		if r, err := calculateRankValue(prefix[0]); err == nil {
			from.Rank = r
			prefix = prefix[1:]
		}
	}
	if len(prefix) > 0 {
		return Move{}, fmt.Errorf("notação inválida: %q não é uma desambiguação de origem", prefix)
	}

	// A captura de peão sempre nomeia a coluna de origem ("exd5"); sem ela a
	// notação está incompleta e não dá para adivinhar de onde o peão veio.
	if capture && piece == Pawn && from.File < 0 {
		return Move{}, fmt.Errorf("captura de peão precisa da coluna de origem, como \"exd5\"")
	}

	if promotion != 0 {
		if piece != Pawn {
			return Move{}, fmt.Errorf("só peões promovem, não %v", piece)
		}
		if rank != 0 && rank != 7 {
			return Move{}, fmt.Errorf("promoção só acontece na primeira ou na última linha")
		}
	}

	return Move{
		Piece:     piece,
		From:      from,
		Capture:   capture,
		Check:     check,
		Checkmate: checkmate,
		Promotion: promotion,
		To: Square{
			File: file,
			Rank: rank,
		},
	}, nil
}

// parseCastle reconhece as grafias usuais do roque, incluindo as que usam o
// dígito zero no lugar da letra O. Devolve CastleNone se não for um roque.
func parseCastle(cmd string) CastleSide {
	switch strings.ToUpper(strings.ReplaceAll(cmd, "0", "O")) {
	case "O-O":
		return CastleKingside
	case "O-O-O":
		return CastleQueenside
	}
	return CastleNone
}

// upper devolve a versão maiúscula de uma letra ASCII.
func upper(charByte byte) byte {
	if charByte >= 'a' && charByte <= 'z' {
		return charByte - 'a' + 'A'
	}
	return charByte
}
