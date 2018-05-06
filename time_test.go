package cstrftime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	tm := time.Date(2000, 2, 1, 16, 30, 22, 123456789, time.UTC)

	cases := []struct {
		name, format, result string
	}{
		{"weekday as locale’s abbreviated name", "%a", "Tue"},
		{"weekday as locale’s full name", "%A", "Tuesday"},
		{"weekday as a decimal number", "%w", "2"},
		{"day of the month as a zero-padded decimal number", "%d", "01"},
		{"month as locale’s abbreviated name", "%b", "Feb"},
		{"month as locale’s full name", "%B", "February"},
		{"month as a zero-padded decimal number", "%m", "02"},
		{"24 hour Hours", "%H", "16"},
		{"12 hour Hours", "%I", "04"},
		{"Minutes", "%M", "30"},
		{"Seconds", "%S", "22"},
		{"seconds with decimals", "%S.%f", "22.123456"},
		{"american date", "%m/%d/%Y", "02/01/2000"},
		{"proper date", "%Y/%m/%d", "2000/02/01"},
		{"dashed date", "%Y-%m-%d", "2000-02-01"},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.result, Format(c.format, tm))
		})
	}
}

func BenchmarkFormat(b *testing.B) {
	tm := time.Date(2000, time.February, 13, 16, 30, 20, 12345, time.UTC)

	cases := []struct {
		name, cFormat, goFormat string
	}{
		{"year", "%Y", "2006"},
		{"date", "%Y/%m/%d", "2006/01/02"},
	}

	for _, c := range cases {
		assert.Equal(b, tm.Format(c.goFormat), Format(c.cFormat, tm), c.name)
	}

	for _, c := range cases {
		for _, native := range []bool{true, false} {
			c := c
			native := native

			b.Run(fmt.Sprintf("%s - native(%t)", c.name, native), func(b *testing.B) {
				var (
					want = tm.Format(c.goFormat)
					got  string
				)

				switch native {
				case true:
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						if got = tm.Format(c.goFormat); want != got {
							b.Fatalf("want: %q got: %q", want, got)
						}
					}

				case false:
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						if got = Format(c.cFormat, tm); want != got {
							b.Fatalf("want: %q got: %q", want, got)
						}
					}
				}
			})
		}
	}
}
