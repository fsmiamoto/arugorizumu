package stairs

import "testing"

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "Base case #1",
		args: args{
			n: 0,
		},
		want: 1,
	},
	{
		name: "Base case #2",
		args: args{
			n: 1,
		},
		want: 1,
	},
	{
		name: "Simple case #1",
		args: args{
			n: 2,
		},
		want: 2,
	},
	{
		name: "Simple case #2",
		args: args{
			n: 5,
		},
		want: 8,
	},
	{
		name: "Big case",
		args: args{
			n: 1_000_000,
		},
		want: 2756670985995446685,
	},
}

func TestClimbStairs(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs(%v) = %v, but want %v", tt.args.n, got, tt.want)
			}
		})
	}
}

func TestClimbStairsConstantSpace(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsConstantSpace(tt.args.n); got != tt.want {
				t.Errorf("climbStairs(%v) = %v, but want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
