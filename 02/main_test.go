package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test",
			args: args{
				input: []string{
					"7 6 4 2 1",
					"1 2 7 8 9",
				},
			},
			want: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafe(t *testing.T) {
	type args struct {
		slice [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				slice: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Safe(tt.args.slice); got != tt.want {
				t.Errorf("Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetterSafe(t *testing.T) {
	type args struct {
		slice [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				slice: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BetterSafe(tt.args.slice); got != tt.want {
				t.Errorf("Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}
