# chess_notation_parser

Biblioteca em Go para interpretar a [notação algébrica](https://pt.wikipedia.org/wiki/Nota%C3%A7%C3%A3o_alg%C3%A9brica_de_xadrez) de lances de xadrez.

## Instalação

```sh
go get github.com/nataferreiradev/chess_notation_parser
```

## Uso

```go
package main

import (
	"fmt"

	chess "github.com/nataferreiradev/chess_notation_parser"
)

func main() {
	move, err := chess.Parse("Bxe4+")
	if err != nil {
		panic(err)
	}

	fmt.Println(move.String()) // Bxe4+
	fmt.Println(move.Piece)    // Bishop
	fmt.Println(move.Check)    // true
}
```

## API

- `chess.Parse(s string) (Move, error)` — interpreta uma jogada.
- `Move` — campos `Piece`, `From`, `To`, `Capture`, `Check`, `Checkmate`.
- `Move.String() string` — reconstrói a notação.
- `Square` — `File` (coluna a=0..h=7) e `Rank` (linha 1=0..8=7); `-1` = não especificado.
- `Piece` — `Pawn`, `Knight`, `Bishop`, `Rook`, `Queen`, `King`.

## CLI de exemplo

```sh
go run ./cmd/chessparse Bxe4+
```

## Testes

```sh
go test ./...
```
