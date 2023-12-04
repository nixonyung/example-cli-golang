package algorithms_test

import (
	"example-cli/pkg/algorithms"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestGCD(t *testing.T) {
	is := is.New(t)

	for _, testcase := range []struct {
		in1   uint
		in2   uint
		want1 uint
	}{
		{in1: 3, in2: 1, want1: 1},
		{in1: 5, in2: 5, want1: 5},
		{in1: 20, in2: 6, want1: 2},
		{in1: 100, in2: 52, want1: 4},
	} {
		t.Run(
			fmt.Sprint(testcase),
			func(_ *testing.T) {
				out1 := algorithms.GCD(testcase.in1, testcase.in2)
				is.Equal(out1, testcase.want1)

				out1_alt := algorithms.GCD(testcase.in1, testcase.in2)
				is.Equal(out1_alt, testcase.want1)
			},
		)
	}
}
