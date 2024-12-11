package main

import (
	"reflect"
	"testing"
)

func TestBlink(t *testing.T) {
	type args struct {
		m M
	}
	tests := []struct {
		name string
		args args
		want M
	}{
		{
			name: "test",
			args: args{
				m: M{
					0:   1,
					1:   1,
					10:  1,
					99:  1,
					999: 1,
				},
			},
			want: M{
				0:       1,
				1:       2,
				9:       2,
				2024:    1,
				2021976: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blink(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlinks(t *testing.T) {
	type args struct {
		m M
		n int
	}
	tests := []struct {
		name string
		args args
		want M
	}{
		{
			name: "test",
			args: args{
				m: M{
					125: 1,
					17:  1,
				},
				n: 6,
			},
			want: M{
				0:          2,
				2:          4,
				3:          1,
				4:          1,
				6:          2,
				7:          1,
				8:          1,
				40:         2,
				48:         2,
				80:         1,
				96:         1,
				2024:       1,
				4048:       1,
				14168:      1,
				2097446912: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blinks(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
