package main

import "testing"

func TestPositions(t *testing.T) {
	type args struct {
		slice [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				slice: [][]rune{
					{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
					{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				},
			},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Positions(tt.args.slice); got != tt.want {
				t.Errorf("Positions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPossibleObstructions(t *testing.T) {
	type args struct {
		slice [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				slice: [][]rune{
					{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
					{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PossibleObstructions(tt.args.slice); got != tt.want {
				t.Errorf("PossibleObstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
