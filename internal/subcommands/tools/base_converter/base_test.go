package base_converter_test

import (
	baseconverter "example-cli/internal/subcommands/tools/base_converter"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestParseBase(t *testing.T) {
	is := is.New(t)

	for _, testcase := range []struct {
		in1   int
		want1 baseconverter.Base
	}{
		{in1: 2, want1: baseconverter.BASE_BINARY},

		{in1: 10, want1: baseconverter.BASE_DECIMAL},

		{in1: 16, want1: baseconverter.BASE_HEXADECIMAL},

		{in1: 0, want1: baseconverter.BASE_UNDEFINED},
		{in1: 3, want1: baseconverter.BASE_UNDEFINED},
		{in1: 123, want1: baseconverter.BASE_UNDEFINED},
		{in1: -1, want1: baseconverter.BASE_UNDEFINED},
	} {
		t.Run(
			fmt.Sprint(testcase),
			func(_ *testing.T) {
				out1 := baseconverter.ParseBase(testcase.in1)
				is.Equal(out1, testcase.want1)
			},
		)
	}
}
