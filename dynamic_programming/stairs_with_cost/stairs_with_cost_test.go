package stairs_with_cost

import (
	"reflect"
	"testing"
)

type args struct {
	n     int
	k     int
	price []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"Base case #1", args{n: 0, k: 2, price: []int{0}}, 0},
	{"Base case #2", args{n: 1, k: 2, price: []int{0, 10, 2}}, 10},
	{"Base case #3", args{n: 2, k: 2, price: []int{0, 10, 2}}, 2},
	{"Simple case #1", args{n: 3, k: 2, price: []int{0, 3, 2, 4}}, 6},
	{"Simple case #2", args{n: 3, k: 2, price: []int{0, 2, 4, 3}}, 5},
	{"Simple case #3", args{n: 8, k: 2, price: []int{0, 3, 2, 4, 6, 1, 1, 5, 3}}, 11},
}

var testsPath = []struct {
	name string
	args args
	want []int
}{
	{"Base case #1", args{n: 0, k: 2, price: []int{0}}, []int{0}},
	{"Base case #2", args{n: 1, k: 2, price: []int{0, 10, 2}}, []int{0, 1}},
	{"Base case #3", args{n: 2, k: 2, price: []int{0, 10, 2}}, []int{0, 2}},
	{"Simple case #1", args{n: 3, k: 2, price: []int{0, 3, 2, 4}}, []int{0, 2, 3}},
	{"Simple case #2", args{n: 3, k: 2, price: []int{0, 2, 4, 3}}, []int{0, 1, 3}},
	{"Simple case #3", args{n: 8, k: 2, price: []int{0, 3, 2, 4, 6, 1, 1, 5, 3}}, []int{0, 2, 3, 5, 6, 8}},
}

func Test_costOfClimbingToStair(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := costOfClimbingToStair(tt.args.n, tt.args.k, tt.args.price); got != tt.want {
				t.Errorf("costOfClimbingToStair(%v,%v,%v) = %v, but want %v",
					tt.args.n, tt.args.k, tt.args.price, got, tt.want)
			}
		})
	}
}
func Test_pathOfClimbingToStair(t *testing.T) {
	for _, tt := range testsPath {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathOfClimbingToStair(tt.args.n, tt.args.k, tt.args.price); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("costOfClimbingToStair(%v,%v,%v) = %v, but want %v",
					tt.args.n, tt.args.k, tt.args.price, got, tt.want)
			}
		})
	}
}
