package model

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestArrMapLen(t *testing.T) {
	anArray := []int{
		1, 2, 3, 4,
	}
	assert.Equal(t, 4, len(anArray))
	val, err := ArrMapLen(reflect.ValueOf(anArray), []reflect.Value{})
	assert.NoError(t, err)
	assert.True(t, val.IsValid())
	assert.Equal(t, 4, int(val.Int()))
}

func TestStrCompare(t *testing.T) {
	a := "A"
	b := "B"
	assert.True(t, strings.Compare(a, b) < 0)

	val, err := StrCompare(a, []reflect.Value{reflect.ValueOf(b)})
	assert.NoError(t, err)
	assert.True(t, val.Int() < 0)
}

func TestStrContains(t *testing.T) {
	a := "ABCDEFG"
	b := "CDE"
	assert.True(t, strings.Contains(a, b))

	val, err := StrContains(a, []reflect.Value{reflect.ValueOf(b)})
	assert.NoError(t, err)
	assert.True(t, val.Bool())
}

func TestStrCount(t *testing.T) {
	a := "BCABCABCABCAB"
	b := "ABC"
	assert.Equal(t, 3, strings.Count(a, b))

	val, err := StrCount(a, []reflect.Value{reflect.ValueOf(b)})
	assert.NoError(t, err)
	assert.Equal(t, 3, int(val.Int()))
}
