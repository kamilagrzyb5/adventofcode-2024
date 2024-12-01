package main

import (
	"reflect"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				s1: []int{3, 4, 2, 1, 3, 3},
				s2: []int{4, 3, 5, 3, 9, 3},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDistance(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "test",
			args: args{
				lines: []string{
					"3   4",
					"4   3",
					"2   5",
					"1   3",
					"3   9",
					"3   3",
				},
			},
			want: []int{
				3, 4, 2, 1, 3, 3,
			},
			want1: []int{
				4, 3, 5, 3, 9, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInput(tt.args.lines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCalculateSimilarity(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				s1: []int{3, 4, 2, 1, 3, 3},
				s2: []int{4, 3, 5, 3, 9, 3},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSimilarity(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CalculateSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
