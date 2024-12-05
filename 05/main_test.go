package main

import (
	"reflect"
	"testing"
)

func TestOrderMap(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want map[int][]int
	}{
		{
			name: "test",
			args: args{
				arr: [][]int{
					{47, 53},
					{97, 13},
					{97, 61},
					{97, 47},
					{75, 29},
					{61, 13},
					{75, 53},
					{29, 13},
					{97, 29},
					{53, 29},
					{61, 53},
					{97, 53},
					{61, 29},
					{47, 13},
					{75, 47},
					{97, 75},
					{47, 61},
					{75, 61},
					{47, 29},
					{75, 13},
					{53, 13},
				},
			},
			want: map[int][]int{
				29: {13},
				47: {53, 13, 61, 29},
				53: {29, 13},
				61: {13, 53, 29},
				75: {29, 53, 47, 61, 13},
				97: {13, 61, 47, 29, 53, 75},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderMap(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCorrects(t *testing.T) {
	type args struct {
		rows [][]int
		m    map[int][]int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		nonWant [][]int
	}{
		{
			name: "test",
			args: args{
				rows: [][]int{
					{75, 47, 61, 53, 29},
					{97, 61, 53, 29, 13},
					{75, 29, 13},
					{75, 97, 47, 61, 53},
					{61, 13, 29},
					{97, 13, 75, 29, 47},
				},
				m: map[int][]int{
					29: {13},
					47: {53, 13, 61, 29},
					53: {29, 13},
					61: {13, 53, 29},
					75: {29, 53, 47, 61, 13},
					97: {13, 61, 47, 29, 53, 75},
				},
			},
			want: [][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
			},
			nonWant: [][]int{
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotNon := Corrects(tt.args.rows, tt.args.m)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Corrects() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(gotNon, tt.nonWant) {
				t.Errorf("Corrects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				[][]int{
					{75, 47, 61, 53, 29},
					{97, 61, 53, 29, 13},
					{75, 29, 13},
				},
			},
			want: 143,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.arr); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFirst(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test",
			args: args{
				s: []string{
					"47|53",
					"97|13",
					"97|61",
					"97|47",
				},
			},
			want: [][]int{
				{47, 53},
				{97, 13},
				{97, 61},
				{97, 47},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFirst(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSecond(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test",
			args: args{
				s: []string{
					"75,47,61,53,29",
					"97,61,53,29,13",
				},
			},
			want: [][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSecond(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNonCorrects(t *testing.T) {
	type args struct {
		rows [][]int
		m    map[int][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test",
			args: args{
				rows: [][]int{
					{75, 97, 47, 61, 53},
					{61, 13, 29},
					{97, 13, 75, 29, 47},
				},
				m: map[int][]int{
					29: {13},
					47: {53, 13, 61, 29},
					53: {29, 13},
					61: {13, 53, 29},
					75: {29, 53, 47, 61, 13},
					97: {13, 61, 47, 29, 53, 75},
				},
			},
			want: [][]int{
				{97, 75, 47, 61, 53},
				{61, 29, 13},
				{97, 75, 47, 29, 13},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonCorrects(tt.args.rows, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonCorrects() = %v, want %v", got, tt.want)
			}
		})
	}
}
