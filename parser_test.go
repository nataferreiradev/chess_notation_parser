package chess

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Move
	}{
		{
			name: "lance de peça",
			in:   "Be4",
			want: Move{Piece: Bishop, From: Square{-1, -1}, To: Square{File: 4, Rank: 3}},
		},
		{
			name: "captura com xeque",
			in:   "Bxe4+",
			want: Move{Piece: Bishop, From: Square{-1, -1}, To: Square{4, 3}, Capture: true, Check: true},
		},
		{
			name: "lance de peão",
			in:   "e4",
			want: Move{Piece: Pawn, From: Square{-1, -1}, To: Square{4, 3}},
		},
		{
			name: "captura de peão com desambiguação e mate",
			in:   "dxe4#",
			want: Move{Piece: Pawn, From: Square{File: 3, Rank: -1}, To: Square{4, 3}, Capture: true, Checkmate: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.in)
			if err != nil {
				t.Fatalf("Parse(%q) erro inesperado: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("Parse(%q) = %+v, esperado %+v", tt.in, got, tt.want)
			}
			if got.String() != tt.in {
				t.Errorf("String() = %q, esperado %q", got.String(), tt.in)
			}
		})
	}
}

func TestParseErros(t *testing.T) {
	for _, in := range []string{"", "e", "i4", "e9"} {
		if _, err := Parse(in); err == nil {
			t.Errorf("Parse(%q) deveria falhar, mas não falhou", in)
		}
	}
}

func TestParseDesambiguacao(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Move
	}{
		{
			name: "desambiguação por coluna",
			in:   "Nbd2",
			want: Move{Piece: Knight, From: Square{File: 1, Rank: -1}, To: Square{File: 3, Rank: 1}},
		},
		{
			name: "desambiguação por linha",
			in:   "R1a3",
			want: Move{Piece: Rook, From: Square{File: -1, Rank: 0}, To: Square{File: 0, Rank: 2}},
		},
		{
			name: "desambiguação por casa inteira",
			in:   "Qh4e1",
			want: Move{Piece: Queen, From: Square{File: 7, Rank: 3}, To: Square{File: 4, Rank: 0}},
		},
		{
			name: "desambiguação por linha com captura",
			in:   "R1xa3",
			want: Move{Piece: Rook, From: Square{File: -1, Rank: 0}, To: Square{File: 0, Rank: 2}, Capture: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.in)
			if err != nil {
				t.Fatalf("Parse(%q) erro inesperado: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("Parse(%q) = %+v, esperado %+v", tt.in, got, tt.want)
			}
			if got.String() != tt.in {
				t.Errorf("String() = %q, esperado %q", got.String(), tt.in)
			}
		})
	}
}

func TestParseRoque(t *testing.T) {
	tests := []struct {
		in   string
		want Move
		str  string
	}{
		{
			in:   "O-O",
			want: Move{Piece: King, Castle: CastleKingside, From: Square{-1, -1}, To: Square{-1, -1}},
			str:  "O-O",
		},
		{
			in:   "O-O-O",
			want: Move{Piece: King, Castle: CastleQueenside, From: Square{-1, -1}, To: Square{-1, -1}},
			str:  "O-O-O",
		},
		{
			in:   "O-O+",
			want: Move{Piece: King, Castle: CastleKingside, From: Square{-1, -1}, To: Square{-1, -1}, Check: true},
			str:  "O-O+",
		},
		{
			in:   "O-O-O#",
			want: Move{Piece: King, Castle: CastleQueenside, From: Square{-1, -1}, To: Square{-1, -1}, Checkmate: true},
			str:  "O-O-O#",
		},
		{
			// grafia com zeros, comum em partidas transcritas à mão
			in:   "0-0",
			want: Move{Piece: King, Castle: CastleKingside, From: Square{-1, -1}, To: Square{-1, -1}},
			str:  "O-O",
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Parse(tt.in)
			if err != nil {
				t.Fatalf("Parse(%q) erro inesperado: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("Parse(%q) = %+v, esperado %+v", tt.in, got, tt.want)
			}
			if got.String() != tt.str {
				t.Errorf("String() = %q, esperado %q", got.String(), tt.str)
			}
		})
	}
}

func TestParsePromocao(t *testing.T) {
	tests := []struct {
		in   string
		want Move
	}{
		{
			in:   "e8=Q",
			want: Move{Piece: Pawn, From: Square{-1, -1}, To: Square{File: 4, Rank: 7}, Promotion: Queen},
		},
		{
			in:   "a1=N",
			want: Move{Piece: Pawn, From: Square{-1, -1}, To: Square{File: 0, Rank: 0}, Promotion: Knight},
		},
		{
			in:   "exd8=R+",
			want: Move{Piece: Pawn, From: Square{File: 4, Rank: -1}, To: Square{File: 3, Rank: 7}, Capture: true, Promotion: Rook, Check: true},
		},
		{
			in:   "b8=B#",
			want: Move{Piece: Pawn, From: Square{-1, -1}, To: Square{File: 1, Rank: 7}, Promotion: Bishop, Checkmate: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Parse(tt.in)
			if err != nil {
				t.Fatalf("Parse(%q) erro inesperado: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("Parse(%q) = %+v, esperado %+v", tt.in, got, tt.want)
			}
			if got.String() != tt.in {
				t.Errorf("String() = %q, esperado %q", got.String(), tt.in)
			}
		})
	}
}

func TestParseAnotacoesSaoToleradas(t *testing.T) {
	base, err := Parse("e4")
	if err != nil {
		t.Fatalf("Parse(\"e4\") erro inesperado: %v", err)
	}

	for _, in := range []string{"e4!", "e4?", "e4!?", "e4!!", "e4??", " e4 "} {
		got, err := Parse(in)
		if err != nil {
			t.Errorf("Parse(%q) erro inesperado: %v", in, err)
			continue
		}
		if got != base {
			t.Errorf("Parse(%q) = %+v, esperado %+v", in, got, base)
		}
	}
}

func TestParseRejeitaLixo(t *testing.T) {
	// Cada entrada é sintaticamente inválida e tem de virar erro, e não um
	// lance com o prefixo silenciosamente descartado.
	tests := []string{
		"Nxxe4",   // dois 'x'
		"Nabce4",  // desambiguação impossível
		"e4=Q",    // promoção fora da última linha
		"Ne8=Q",   // peça maior não promove
		"e8=K",    // não se promove a rei
		"e8=P",    // nem a peão
		"e8=",     // promoção sem peça
		"e8=QQ",   // promoção com lixo
		"O-O-O-O", // roque inexistente
		"xe4",     // captura sem origem nem peça
		"4e",      // invertido
	}

	for _, in := range tests {
		if got, err := Parse(in); err == nil {
			t.Errorf("Parse(%q) deveria falhar, devolveu %+v", in, got)
		}
	}
}

func TestParseNotacaoLongaViraDesambiguacaoCompleta(t *testing.T) {
	// "Ne4e5" não é lixo: é notação longa. O parser preenche a origem inteira
	// e deixa para o motor decidir que nenhum cavalo vai de e4 a e5.
	got, err := Parse("Ne4e5")
	if err != nil {
		t.Fatalf("Parse(\"Ne4e5\") erro inesperado: %v", err)
	}
	want := Move{Piece: Knight, From: Square{File: 4, Rank: 3}, To: Square{File: 4, Rank: 4}}
	if got != want {
		t.Errorf("Parse(\"Ne4e5\") = %+v, esperado %+v", got, want)
	}
}
