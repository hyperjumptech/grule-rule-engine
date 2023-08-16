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

package antlr

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func unquoteString(theStr string) (string, error) {
	strLen := len(theStr)
	if strLen < 2 {
		return "", strconv.ErrSyntax
	}
	quote := theStr[0]
	if quote != theStr[strLen-1] {
		return "", strconv.ErrSyntax
	}
	theStr = theStr[1 : strLen-1]

	if quote != '"' && quote != '\'' {
		return "", strconv.ErrSyntax
	}

	if !contains(theStr, '\\') && !contains(theStr, quote) && utf8.ValidString(theStr) {
		return theStr, nil
	}

	var runeTmp [utf8.UTFMax]byte
	buf := make([]byte, 0, 3*len(theStr)/2)
	for len(theStr) > 0 {
		theRune, multibyte, ss, err := strconv.UnquoteChar(theStr, quote)
		if err != nil {
			return "", err
		}
		theStr = ss
		if theRune < utf8.RuneSelf || !multibyte {
			buf = append(buf, byte(theRune))
		} else {
			n := utf8.EncodeRune(runeTmp[:], theRune)
			buf = append(buf, runeTmp[:n]...)
		}
	}

	return string(buf), nil
}

func contains(s string, c byte) bool {
	return strings.IndexByte(s, c) != -1
}
