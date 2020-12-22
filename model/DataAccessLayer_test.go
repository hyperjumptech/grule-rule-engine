//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

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

func TestStrHasPrefix(t *testing.T) {
	a := "abigbrownfox"
	b := "abig"
	c := "big"

	val, err := StrHasPrefix(a, []reflect.Value{reflect.ValueOf(b)})
	assert.NoError(t, err)
	assert.True(t, val.Bool())

	val, err = StrHasPrefix(a, []reflect.Value{reflect.ValueOf(c)})
	assert.NoError(t, err)
	assert.False(t, val.Bool())
}

func TestStrHasSuffix(t *testing.T) {
	a := "abigbrownfox"
	b := "fox"
	c := "big"

	val, err := StrHasSuffix(a, []reflect.Value{reflect.ValueOf(b)})
	assert.NoError(t, err)
	assert.True(t, val.Bool())

	val, err = StrHasSuffix(a, []reflect.Value{reflect.ValueOf(c)})
	assert.NoError(t, err)
	assert.False(t, val.Bool())
}

func TestStrIndex(t *testing.T) {
	a := "abigfoxbrownfox"
	b := "fox"

	val, err := StrIndex(a, []reflect.Value{reflect.ValueOf(b)})
	assert.NoError(t, err)
	assert.Equal(t, 4, int(val.Int()))
}

func TestStrLastIndex(t *testing.T) {
	a := "abigfoxbrownfox"
	b := "fox"

	val, err := StrLastIndex(a, []reflect.Value{reflect.ValueOf(b)})
	assert.NoError(t, err)
	assert.Equal(t, 12, int(val.Int()))
}

func TestStrLen(t *testing.T) {
	a := "abigfoxbrownfox"

	val, err := StrLen(a, []reflect.Value{})
	assert.NoError(t, err)
	assert.Equal(t, len(a), int(val.Int()))
}

func TestStrRepeat(t *testing.T) {
	a := "foxer"

	val, err := StrRepeat(a, []reflect.Value{reflect.ValueOf(5)})
	assert.NoError(t, err)
	assert.Equal(t, strings.Repeat(a, 5), val.String())
}

func TestStrReplace(t *testing.T) {
	a := "aBigBadFoxJumpsOverALazyFox"
	b := "aBigBadWolfJumpsOverALazyWolf"
	val, err := StrReplace(a, []reflect.Value{reflect.ValueOf("Fox"), reflect.ValueOf("Wolf")})
	assert.NoError(t, err)
	assert.Equal(t, b, val.String())
}

func TestStrSplit(t *testing.T) {
	a := "Big,Bad,Ugly"
	val, err := StrSplit(a, []reflect.Value{reflect.ValueOf(",")})
	assert.NoError(t, err)
	assert.Equal(t, reflect.Slice, val.Kind())
	assert.Equal(t, 3, val.Len())
	assert.Equal(t, "Big", val.Index(0).String())
	assert.Equal(t, "Bad", val.Index(1).String())
	assert.Equal(t, "Ugly", val.Index(2).String())
}

func TestStrToLower(t *testing.T) {
	a := "Big,Bad,Ugly"
	val, err := StrToLower(a, []reflect.Value{})
	assert.NoError(t, err)
	assert.Equal(t, "big,bad,ugly", val.String())
}

func TestStrToUpper(t *testing.T) {
	a := "Big,Bad,Ugly"
	val, err := StrToUpper(a, []reflect.Value{})
	assert.NoError(t, err)
	assert.Equal(t, "BIG,BAD,UGLY", val.String())
}

func TestStrTrim(t *testing.T) {
	a := "   Big,Bad,Ugly   "
	val, err := StrTrim(a, []reflect.Value{})
	assert.NoError(t, err)
	assert.Equal(t, strings.TrimSpace(a), val.String())
}
