package algorithms_test

import (
	"example-cli/pkg/algorithms"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestPrimes(t *testing.T) {
	is := is.New(t)

	for _, testcase := range []struct {
		in1   int
		want1 []int
	}{
		{in1: 2, want1: []int{2}},
		{in1: 3, want1: []int{2, 3}},
		{in1: 10, want1: []int{2, 3, 5, 7}},
		{in1: 100, want1: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	} {
		t.Run(
			fmt.Sprint(testcase),
			func(_ *testing.T) {
				out1 := algorithms.PrimesUpTo(testcase.in1)
				is.Equal(out1, testcase.want1)
			},
		)
	}
}
