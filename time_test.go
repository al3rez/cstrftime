package cstrftime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	tm := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)

	t.Run("weekday as locale’s abbreviated name", func(t *testing.T) {
		assert.Equal(t, "Tue", Format("%a", tm))
	})

	t.Run("weekday as locale’s full name", func(t *testing.T) {
		assert.Equal(t, "Tuesday", Format("%A", tm))
	})

	t.Run("weekday as a decimal number", func(t *testing.T) {
		assert.Equal(t, "2", Format("%w", tm))
	})

	t.Run("day of the month as a zero-padded decimal number", func(t *testing.T) {
		assert.Equal(t, "01", Format("%d", tm))
	})

	t.Run("month as locale’s abbreviated name", func(t *testing.T) {
		assert.Equal(t, "Feb", Format("%b", tm))
	})

	t.Run("month as locale’s full name", func(t *testing.T) {
		assert.Equal(t, "February", Format("%B", tm))
	})

	t.Run("month as a zero-padded decimal number", func(t *testing.T) {
		assert.Equal(t, "02", Format("%m", tm))
	})
}
