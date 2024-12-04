package main

import "testing"

func TestXMAS(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				arr: []string{
					"MMMSXXMASM",
					"MSAMXMSMSA",
					"AMXSXMAAMM",
					"MSAMASMSMX",
					"XMASAMXAMM",
					"XXAMMXXAMA",
					"SMSMSASXSS",
					"SAXAMASAAA",
					"MAMMMXMMMM",
					"MXMXAXMASX",
				},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XMAS(tt.args.arr); got != tt.want {
				t.Errorf("XMAS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMAS(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				arr: []string{
					".M.S......",
					"..A..MSMS.",
					".M.S.MAA..",
					"..A.ASMSM.",
					".M.S.M....",
					"..........",
					"S.S.S.S.S.",
					".A.A.A.A..",
					"M.M.M.M.M.",
					"..........",
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MAS(tt.args.arr); got != tt.want {
				t.Errorf("MAS() = %v, want %v", got, tt.want)
			}
		})
	}
}
