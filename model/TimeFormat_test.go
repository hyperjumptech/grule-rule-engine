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
	"testing"
	"time"
)

/*
ANSIC       = "Mon Jan _2 15:04:05 2006"
UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
RFC822      = "02 Jan 06 15:04 MST"
RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
RFC3339     = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen     = "3:04PM"
// Handy time stamps.
Stamp      = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"
*/

type TimeTestData struct {
	Layout string
	Date   string
	Valid  bool
}

func TestIsDateFormatValid(t *testing.T) {
	testData := []*TimeTestData{
		{Layout: "notexist", Date: "02 Jan 06 15:04 MST", Valid: false},
		{Layout: time.ANSIC, Date: "Mon Jan 2 15:04:05 2006", Valid: true},
		{Layout: time.ANSIC, Date: "Mon Jan 02 15:04:05 2006", Valid: true},
		{Layout: time.ANSIC, Date: "Mon Jan 22 15:04:05 2006", Valid: true},
		{Layout: time.ANSIC, Date: "Mon Jan 22 15:04:05 06", Valid: false},
		// todo add more format test here
	}
	for _, td := range testData {
		v := IsDateFormatValid(td.Layout, td.Date)
		if v != td.Valid {
			t.Logf("layout '%s' expect '%s' to %v but %v", td.Layout, td.Date, td.Valid, !td.Valid)
			t.Fail()
		}
	}
}
