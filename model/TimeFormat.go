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
	"regexp"
	"time"
)

const (
	// ANSIC regex to validate ANSIC date string
	ANSIC = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-3]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2} [0-9]{4}$`
	// UnixDate regex to validate UnixDate date string
	UnixDate = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-3]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2} MST [0-9]{4}$`
	// RubyDate regex to validate RubyDate date string
	RubyDate = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-3][0-9] [0-9]{2}:[0-9]{2}:[0-9]{2} \-?[0-9]{4} [0-9]{4}$`
	// RFC822 regex to validate RFC822 date string
	RFC822 = `^[0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{2} [0-9]{2}:[0-9]{2} MST$`
	// RFC822Z regex to validate RFC822Z date string
	RFC822Z = `^[0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{2} [0-9]{2}:[0-9]{2} \-?[0-9]{4}$` // RFC822 with numeric zone
	// RFC850 regex to validate RFC850 date string
	RFC850 = `^(Monday|Tuesday|Wednesday|Thursday|Friday|Saturday|Sunday), [0-9]{2}-(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2} MST$`
	// RFC1123 regex to validate RFC1123 date string
	RFC1123 = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun), [0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{4} [0-9]{2}:[0-9]{2}:[0-9]{2} MST$`
	// RFC1123Z regex to validate RFC1123Z date string
	RFC1123Z = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun), [0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{4} [0-9]{2}:[0-9]{2}:[0-9]{2} \-?[0-9]{4}$` // RFC1123 with numeric zone
	// RFC3339 regex to validate RFC3339 date string
	RFC3339 = `^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z[0-9]{2}:[0-9]{2}$`
	// RFC3339Nano regex to validate RFC3339Nano date string
	RFC3339Nano = `^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{2}Z[0-9]{2}:[0-9]{2}$`
	// Kitchen regex to validate Kitchen date string
	Kitchen = `^[0-1]?[0-9]:[0-9]{2}(AM|PM)$`
	// Stamp regex to validate Stamp date string
	Stamp = `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2}$`
	// StampMilli regex to validate StampMilli date string
	StampMilli = `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{3}$`
	// StampMicro regex to validate StampMicro date string
	StampMicro = `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{6}$`
	// StampNano regex to validate StampNano date string
	StampNano = `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{9}$`
)

var (
	// TimeValidatorMap map from time format to their validator regex pattern
	TimeValidatorMap = map[string]string{
		time.ANSIC:       ANSIC,
		time.UnixDate:    UnixDate,
		time.RubyDate:    RubyDate,
		time.RFC822:      RFC822,
		time.RFC822Z:     RFC822Z,
		time.RFC850:      RFC850,
		time.RFC1123:     RFC1123,
		time.RFC1123Z:    RFC1123Z,
		time.RFC3339:     RFC3339,
		time.RFC3339Nano: RFC3339Nano,
		time.Kitchen:     Kitchen,
		time.Stamp:       Stamp,
		time.StampMilli:  StampMilli,
		time.StampMicro:  StampMicro,
		time.StampNano:   StampNano,
	}
)

// IsDateFormatValid validate if the supplied date string is compatible with specified format.
// the format should come from standard time format (eg. time.RFC3339, time.ANSIC, time.ANSIC, time.RFC850, etc)
func IsDateFormatValid(layout, date string) bool {
	if pattern, ok := TimeValidatorMap[layout]; ok {
		m, err := regexp.MatchString(pattern, date)
		if err != nil || !m {
			return false
		}
		return true
	}
	return false
}
