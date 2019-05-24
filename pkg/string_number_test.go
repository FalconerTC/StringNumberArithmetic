package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		a := StringNumber{value: "2"}
		b := StringNumber{value: "2"}
		a.Add(b)
		assert.Equal(t, a.value, "4")
	})

	t.Run("first is smaller", func(t *testing.T) {
		a := StringNumber{value: "9"}
		b := StringNumber{value: "214"}
		a.Add(b)
		assert.Equal(t, a.value, "223")
	})

	t.Run("second is smaller", func(t *testing.T) {
		a := StringNumber{value: "356"}
		b := StringNumber{value: "21"}
		a.Add(b)
		assert.Equal(t, a.value, "377")
	})

	t.Run("empty strings", func(t *testing.T) {
		a := StringNumber{value: ""}
		b := StringNumber{value: ""}
		a.Add(b)
		assert.Equal(t, a.value, "0")
	})

	t.Run("one empty string", func(t *testing.T) {
		a := StringNumber{value: ""}
		b := StringNumber{value: "6"}
		a.Add(b)
		assert.Equal(t, a.value, "6")
	})

	t.Run("string with non-numbers", func(t *testing.T) {
		a := StringNumber{value: "a9"}
		b := StringNumber{value: "13"}
		a.Add(b)
		assert.Equal(t, a.value, "22")
	})

	t.Run("really big strings", func(t *testing.T) {
		a := StringNumber{value: "1111111111111111111111111111111111111111"}
		b := StringNumber{value: "2222222222222222222222222222222222222222"}
		a.Add(b)
		assert.Equal(t, a.value, "3333333333333333333333333333333333333333")
	})
}

func Test_Multiply(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		a := StringNumber{value: "2"}
		b := StringNumber{value: "3"}
		a.Multiply(b)
		assert.Equal(t, a.value, "6")
	})

	t.Run("first is smaller", func(t *testing.T) {
		a := StringNumber{value: "9"}
		b := StringNumber{value: "214"}
		a.Multiply(b)
		assert.Equal(t, a.value, "1926")
	})

	t.Run("second is smaller", func(t *testing.T) {
		a := StringNumber{value: "356"}
		b := StringNumber{value: "21"}
		a.Multiply(b)
		assert.Equal(t, a.value, "7476")
	})

	t.Run("empty strings", func(t *testing.T) {
		a := StringNumber{value: ""}
		b := StringNumber{value: ""}
		a.Multiply(b)
		assert.Equal(t, a.value, "0")
	})

	t.Run("one empty string", func(t *testing.T) {
		a := StringNumber{value: ""}
		b := StringNumber{value: "6"}
		a.Multiply(b)
		assert.Equal(t, a.value, "0")
	})

	t.Run("string with non-numbers", func(t *testing.T) {
		a := StringNumber{value: "a9"}
		b := StringNumber{value: "13"}
		a.Multiply(b)
		assert.Equal(t, a.value, "117")
	})

	t.Run("really big strings", func(t *testing.T) {
		a := StringNumber{value: "1111111111111111111111111111111111111111"}
		b := StringNumber{value: "2222222222222222222222222222222222222222"}
		a.Multiply(b)
		assert.Equal(t, a.value, "2469135802469135802469135802469135802468641975308641975308641975308641975308642")
	})
}
