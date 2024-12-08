package main

import (
	"reflect"
	"testing"
)

func TestMakePoints(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			name: "test",
			args: args{
				s: []string{
					"....",
					".A..",
					"..A.",
					"....",
				},
			},
			want: []Point{
				{i: 0, j: 0, sign: '.'},
				{i: 0, j: 1, sign: '.'},
				{i: 0, j: 2, sign: '.'},
				{i: 0, j: 3, sign: '.'},
				{i: 1, j: 0, sign: '.'},
				{i: 1, j: 1, sign: 'A'},
				{i: 1, j: 2, sign: '.'},
				{i: 1, j: 3, sign: '.'},
				{i: 2, j: 0, sign: '.'},
				{i: 2, j: 1, sign: '.'},
				{i: 2, j: 2, sign: 'A'},
				{i: 2, j: 3, sign: '.'},
				{i: 3, j: 0, sign: '.'},
				{i: 3, j: 1, sign: '.'},
				{i: 3, j: 2, sign: '.'},
				{i: 3, j: 3, sign: '.'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakePoints(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Points() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAntinodes(t *testing.T) {
	type args struct {
		points Points
	}
	tests := []struct {
		name string
		args args
		want Points
	}{
		{
			name: "test",
			args: args{
				points: Points{
					{i: 0, j: 0, sign: '.'},
					{i: 0, j: 1, sign: '.'},
					{i: 0, j: 2, sign: '.'},
					{i: 0, j: 3, sign: '.'},
					{i: 1, j: 0, sign: '.'},
					{i: 1, j: 1, sign: 'A'},
					{i: 1, j: 2, sign: '.'},
					{i: 1, j: 3, sign: '.'},
					{i: 2, j: 0, sign: '.'},
					{i: 2, j: 1, sign: '.'},
					{i: 2, j: 2, sign: 'A'},
					{i: 2, j: 3, sign: '.'},
					{i: 3, j: 0, sign: '.'},
					{i: 3, j: 1, sign: '.'},
					{i: 3, j: 2, sign: '.'},
					{i: 3, j: 3, sign: '.'},
				},
			},
			want: Points{
				{i: 0, j: 0, sign: '.', antinode: true},
				{i: 0, j: 1, sign: '.'},
				{i: 0, j: 2, sign: '.'},
				{i: 0, j: 3, sign: '.'},
				{i: 1, j: 0, sign: '.'},
				{i: 1, j: 1, sign: 'A'},
				{i: 1, j: 2, sign: '.'},
				{i: 1, j: 3, sign: '.'},
				{i: 2, j: 0, sign: '.'},
				{i: 2, j: 1, sign: '.'},
				{i: 2, j: 2, sign: 'A'},
				{i: 2, j: 3, sign: '.'},
				{i: 3, j: 0, sign: '.'},
				{i: 3, j: 1, sign: '.'},
				{i: 3, j: 2, sign: '.'},
				{i: 3, j: 3, sign: '.', antinode: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Antinodes(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Antinodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtraAntinodes(t *testing.T) {
	type args struct {
		points Points
	}
	tests := []struct {
		name string
		args args
		want Points
	}{
		{
			name: "test",
			args: args{
				points: Points{
					{i: 0, j: 0, sign: '.'},
					{i: 0, j: 1, sign: '.'},
					{i: 0, j: 2, sign: '.'},
					{i: 0, j: 3, sign: '.'},
					{i: 0, j: 4, sign: '.'},
					{i: 1, j: 0, sign: '.'},
					{i: 1, j: 1, sign: 'A'},
					{i: 1, j: 2, sign: '.'},
					{i: 1, j: 3, sign: '.'},
					{i: 1, j: 4, sign: '.'},
					{i: 2, j: 0, sign: '.'},
					{i: 2, j: 1, sign: '.'},
					{i: 2, j: 2, sign: 'A'},
					{i: 2, j: 3, sign: '.'},
					{i: 2, j: 4, sign: '.'},
					{i: 3, j: 0, sign: '.'},
					{i: 3, j: 1, sign: '.'},
					{i: 3, j: 2, sign: '.'},
					{i: 3, j: 3, sign: '.'},
					{i: 3, j: 4, sign: '.'},
					{i: 4, j: 0, sign: '.'},
					{i: 4, j: 1, sign: '.'},
					{i: 4, j: 2, sign: '.'},
					{i: 4, j: 3, sign: '.'},
					{i: 4, j: 4, sign: '.'},
				},
			},
			want: Points{
				{i: 0, j: 0, sign: '.', antinode: true},
				{i: 0, j: 1, sign: '.'},
				{i: 0, j: 2, sign: '.'},
				{i: 0, j: 3, sign: '.'},
				{i: 0, j: 4, sign: '.'},
				{i: 1, j: 0, sign: '.'},
				{i: 1, j: 1, sign: 'A', antinode: true},
				{i: 1, j: 2, sign: '.'},
				{i: 1, j: 3, sign: '.'},
				{i: 1, j: 4, sign: '.'},
				{i: 2, j: 0, sign: '.'},
				{i: 2, j: 1, sign: '.'},
				{i: 2, j: 2, sign: 'A', antinode: true},
				{i: 2, j: 3, sign: '.'},
				{i: 2, j: 4, sign: '.'},
				{i: 3, j: 0, sign: '.'},
				{i: 3, j: 1, sign: '.'},
				{i: 3, j: 2, sign: '.'},
				{i: 3, j: 3, sign: '.', antinode: true},
				{i: 3, j: 4, sign: '.'},
				{i: 4, j: 0, sign: '.'},
				{i: 4, j: 1, sign: '.'},
				{i: 4, j: 2, sign: '.'},
				{i: 4, j: 3, sign: '.'},
				{i: 4, j: 4, sign: '.', antinode: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtraAntinodes(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtraAntinodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
