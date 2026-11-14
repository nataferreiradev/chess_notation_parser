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
- `Move` — campos `Piece`, `From`, `To`, `Capture`, `Check`, `Checkmate`,
  `Castle`, `Promotion`.
- `Move.String() string` — reconstrói a notação.
- `Square` — `File` (coluna a=0..h=7) e `Rank` (linha 1=0..8=7); `-1` = não especificado.
- `Piece` — `Pawn`, `Knight`, `Bishop`, `Rook`, `Queen`, `King`.
- `CastleSide` — `CastleNone`, `CastleKingside`, `CastleQueenside`.

## O que a notação cobre

| Forma | Exemplo | Resultado |
|---|---|---|
| lance simples | `e4`, `Nf3` | `To` preenchido, `From` `{-1,-1}` |
| captura | `exd5`, `Bxe4` | `Capture: true` |
| desambiguação por coluna | `Nbd2` | `From.File` preenchido |
| desambiguação por linha | `R1a3` | `From.Rank` preenchido |
| notação longa | `Qh4e1` | `From` inteiro preenchido |
| roque | `O-O`, `O-O-O`, `0-0` | `Castle`; `From` e `To` ficam `{-1,-1}` |
| promoção | `e8=Q`, `exd8=N+` | `Promotion` preenchido |
| xeque e mate | `Bxe4+`, `Qh7#` | `Check` / `Checkmate` |
| anotações | `e4!?` | toleradas e descartadas |

O parser é só sintático: não conhece o tabuleiro, então não sabe se a peça
existe nem se alcança o destino. Num roque as casas dependem da cor de quem
joga, que a notação não carrega — por isso `From` e `To` vêm vazios e cabe a
quem tem o tabuleiro preenchê-los.

## CLI de exemplo

```sh
go run ./cmd/chessparse Bxe4+
```

## Testes

```sh
go test ./...
```
