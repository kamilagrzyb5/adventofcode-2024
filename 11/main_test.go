package main

import (
	"reflect"
	"testing"
)

func TestBlink(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				arr: []int{0, 1, 10, 99, 999},
			},
			want: []int{1, 2024, 1, 0, 9, 9, 2021976},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blink(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlinks(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				arr: []int{125, 17},
				n:   6,
			},
			want: []int{
				2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blinks(tt.args.arr, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
