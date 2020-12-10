package model

import (
	"regexp"
	"time"
)

const (
	ANSIC       = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-3]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2} [0-9]{4}$`
	UnixDate    = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-3]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2} MST [0-9]{4}$`
	RubyDate    = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-3][0-9] [0-9]{2}:[0-9]{2}:[0-9]{2} \-?[0-9]{4} [0-9]{4}$`
	RFC822      = `^[0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{2} [0-9]{2}:[0-9]{2} MST$`
	RFC822Z     = `^[0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{2} [0-9]{2}:[0-9]{2} \-?[0-9]{4}$` // RFC822 with numeric zone
	RFC850      = `^(Monday|Tuesday|Wednesday|Thursday|Friday|Saturday|Sunday), [0-9]{2}-(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2} MST$`
	RFC1123     = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun), [0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{4} [0-9]{2}:[0-9]{2}:[0-9]{2} MST$`
	RFC1123Z    = `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun), [0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{4} [0-9]{2}:[0-9]{2}:[0-9]{2} \-?[0-9]{4}$` // RFC1123 with numeric zone
	RFC3339     = `^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z[0-9]{2}:[0-9]{2}$`
	RFC3339Nano = `^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{2}Z[0-9]{2}:[0-9]{2}$`
	Kitchen     = `^[0-1]?[0-9]:[0-9]{2}(AM|PM)$`
	// Handy time stamps.
	Stamp      = `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2}$`
	StampMilli = `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{3}$`
	StampMicro = `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{6}$`
	StampNano  = `^(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]?[0-9] [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{9}$`
)

var (
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
