package main

import (
	"reflect"
	"testing"
)

func TestEvals(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				numbers: []int{
					10, 19,
				},
			},
			want: []int{29, 190},
		},
		{
			name: "2",
			args: args{
				numbers: []int{
					81, 40, 27,
				},
			},
			want: []int{148, 3267, 3267, 87480},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Evals(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAble(t *testing.T) {
	type args struct {
		result int
		arr    []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				result: 190,
				arr:    []int{10, 19},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				result: 3267,
				arr:    []int{81, 40, 27},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				result: 83,
				arr:    []int{17, 5},
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				result: 161011,
				arr:    []int{16, 10, 13},
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				result: 292,
				arr:    []int{11, 6, 16, 20},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Able(tt.args.result, tt.args.arr, Evals); got != tt.want {
				t.Errorf("Able() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
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
				arr: [][]int{
					{190, 10, 19},
					{3267, 81, 40, 27},
					{83, 17, 5},
					{156, 15, 6},
					{7290, 6, 8, 6, 15},
					{161011, 16, 10, 13},
					{192, 17, 8, 14},
					{21037, 9, 7, 18, 13},
					{292, 11, 6, 16, 20},
				},
			},
			want: 3749,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.arr, Evals); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
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
					"190: 10 19",
					"3267: 81 40 27",
				},
			},
			want: [][]int{
				{190, 10, 19},
				{3267, 81, 40, 27},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concatenation(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				a: 193,
				b: 519,
			},
			want: 193519,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatenation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("concatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtraEvals(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				numbers: []int{10, 19},
			},
			want: []int{29, 190, 1019},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtraEvals(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtraEvals() = %v, want %v", got, tt.want)
			}
		})
	}
}
