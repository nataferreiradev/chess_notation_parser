package chess

// Piece representa uma peça de xadrez pela letra usada na notação algébrica.
type Piece byte

const (
	Pawn   Piece = 'P'
	Knight Piece = 'N'
	Bishop Piece = 'B'
	Rook   Piece = 'R'
	Queen  Piece = 'Q'
	King   Piece = 'K'
)

// String devolve o nome da peça em inglês, ou string vazia se desconhecida.
func (p Piece) String() string {
	switch p {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	}
	return ""
}

// decidePiece interpreta um byte como peça; qualquer valor que não seja uma
// peça maior é tratado como peão.
func decidePiece(charByte byte) Piece {
	switch Piece(charByte) {
	case King, Queen, Bishop, Rook, Knight:
		return Piece(charByte)
	default:
		return Pawn
	}
}
