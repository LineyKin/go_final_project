package next_date

import (
	"regexp"
	"strings"
)

func checkRepeatFormat(repeat, repeatType string) bool {
	switch repeatType {
	case yearType:
		return chekYearType(repeat)
	case dayType:
		return checkDayType(repeat)
	default:
		return false
	}
}

func checkDayType(repeat string) bool {
	repeat = strings.TrimSpace(repeat)
	repeat = strings.ToLower(repeat)
	dayRuleRegExp := `^d{1}\s+([1-3]{1}\d{2}|[1-9]{1}\d?|400)$`
	matched, err := regexp.MatchString(dayRuleRegExp, repeat)

	if err != nil {
		panic(err)
	}

	return matched
}

func chekYearType(repeat string) bool {
	repeat = strings.TrimSpace(repeat)
	repeat = strings.ToLower(repeat)

	return repeat == "y"
}
