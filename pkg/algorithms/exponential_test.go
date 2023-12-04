package algorithms_test

import (
	"example-cli/pkg/algorithms"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestExponential(t *testing.T) {
	is := is.New(t)

	for _, testcase := range []struct {
		in1   int
		in2   int
		want1 int
	}{
		{in1: 2, in2: 10, want1: 1024},
		{in1: 3, in2: 6, want1: 729},
	} {
		t.Run(
			fmt.Sprint(testcase),
			func(_ *testing.T) {
				out1 := algorithms.Exponential(testcase.in1, testcase.in2)
				is.Equal(out1, testcase.want1)
			},
		)
	}
}
