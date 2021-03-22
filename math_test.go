package go_utils

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试Add方法",
			args: args{nums: []int{10, 20, 30}},
			want: 60,
		},
		{
			name: "测试Add方法",
			args: args{nums: []int{10, 20, -30}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.nums...); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试多个数相乘",
			args: args{ nums: []int{10, 1, 20} },
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Times(tt.args.nums...); got != tt.want {
				t.Errorf("Times() = %v, want %v", got, tt.want)
			}
		})
	}
}
