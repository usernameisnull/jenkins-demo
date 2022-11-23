package main

import "testing"

func Test_generateRandomNum(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "succ",
			args: args{
				min: 10,
				max: 11,
			},
			want: 10,
		},
		{
			name: "succ",
			args: args{
				min: 11,
				max: 12,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomNum(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("generateRandomNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
