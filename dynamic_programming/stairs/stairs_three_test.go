package stairs

import "testing"

var tests3steps = []struct {
	name string
	args args
	want int
}{
	{
		name: "Simple case #1",
		args: args{
			n: 3,
		},
		want: 4,
	},
	{
		name: "Simple case #2",
		args: args{
			n: 5,
		},
		want: 13,
	},
}

func TestClimbStairsThreeSteps(t *testing.T) {
	for _, tt := range tests3steps {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsThreeSteps(tt.args.n); got != tt.want {
				t.Errorf("climbStairs(%v) = %v, but want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
