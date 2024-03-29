package importer

// thanks random guy on stackoverflow

import (
	"strings"
	"time"
)

var dayOrdinals = map[string]string{ // map[ordinal]cardinal
	"1st": "1", "2nd": "2", "3rd": "3", "4th": "4", "5th": "5",
	"6th": "6", "7th": "7", "8th": "8", "9th": "9", "10th": "10",
	"11th": "11", "12th": "12", "13th": "13", "14th": "14", "15th": "15",
	"16th": "16", "17th": "17", "18th": "18", "19th": "19", "20th": "20",
	"21st": "21", "22nd": "22", "23rd": "23", "24th": "24", "25th": "25",
	"26th": "26", "27th": "27", "28th": "28", "29th": "29", "30th": "30",
	"31st": "31",
}

// parseOrdinalDate parses a string time value using an ordinary package time layout.
// Before parsing, an ordinal day, [1st, 31st], is converted to a cardinal day, [1, 31].
// For example, "1st August 2017" is converted to "1 August 2017" before parsing, and
// "August 1st, 2017" is converted to "August 1, 2017" before parsing.
func parseOrdinalDate(layout, value string) (time.Time, error) {
	const ( // day number
		cardMinLen = len("1")
		cardMaxLen = len("31")
		ordSfxLen  = len("th")
		ordMinLen  = cardMinLen + ordSfxLen
	)

	for k := 0; k < len(value)-ordMinLen; {
		// i number start
		for ; k < len(value) && (value[k] > '9' || value[k] < '0'); k++ {
		}
		i := k
		// j cardinal end
		for ; k < len(value) && (value[k] <= '9' && value[k] >= '0'); k++ {
		}
		j := k
		if j-i > cardMaxLen || j-i < cardMinLen {
			continue
		}
		// k ordinal end
		// ASCII Latin (uppercase | 0x20) = lowercase
		for ; k < len(value) && (value[k]|0x20 >= 'a' && value[k]|0x20 <= 'z'); k++ {
		}
		if k-j != ordSfxLen {
			continue
		}

		// day ordinal to cardinal
		for ; i < j-1 && (value[i] == '0'); i++ {
		}
		o := strings.ToLower(value[i:k])
		c, ok := dayOrdinals[o]
		if ok {
			value = value[:i] + c + value[k:]
			break
		}
	}

	return time.Parse(layout, value)
}
