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
