// Command chessparse interpreta uma jogada em notação algébrica passada como
// argumento e imprime os detalhes do lance.
//
//	chessparse Bxe4+
package main

import (
	"fmt"
	"os"

	chess "github.com/nataferreiradev/chess_notation_parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: chessparse <jogada>")
		os.Exit(2)
	}

	move, err := chess.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}

	fmt.Println("notação:", move.String())
	fmt.Println("peça:   ", move.Piece)
	fmt.Println("destino:", move.To)
	fmt.Println("captura:", move.Capture)
	fmt.Println("xeque:  ", move.Check)
	fmt.Println("mate:   ", move.Checkmate)
}
