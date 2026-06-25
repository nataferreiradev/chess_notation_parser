package chess

// Square representa uma casa do tabuleiro.
//
// File é a coluna: a=0 .. h=7. Rank é a linha: 1=0 .. 8=7.
// O valor -1 indica componente não especificado (por exemplo, a origem de um
// lance sem desambiguação).
type Square struct {
	File int
	Rank int
}

// calculateFileValue converte a letra da coluna ('a'..'h') em índice 0..7.
func calculateFileValue(charByte byte) (int, error) {
	value := int(charByte - 'a')
	if value < 0 || value > 7 {
		return -1, errInvalidFile
	}
	return value, nil
}

// calculateRankValue converte o dígito da linha ('1'..'8') em índice 0..7.
func calculateRankValue(charByte byte) (int, error) {
	value := int(charByte - '1')
	if value < 0 || value > 7 {
		return -1, errInvalidRank
	}
	return value, nil
}
