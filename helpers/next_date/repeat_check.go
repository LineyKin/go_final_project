package next_date

import (
	"fmt"
	"regexp"
	"strings"
)

func checkRepeatFormat(repeat, repeatType string) bool {
	switch repeatType {
	case yearType:
		return chekYearType(repeat)
	case dayType:
		return checkDayType(repeat)
	case weekType:
		return checkWeekType(repeat)
	default:
		return false
	}
}

func isUnique(arr []string) bool {
	fmt.Println(arr)
	valMap := make(map[string]bool)

	for _, day := range arr {
		_, ok := valMap[day]
		if !ok {
			valMap[day] = true
		}
	}

	return len(valMap) == len(arr)
}

func checkWeekType(repeat string) bool {
	repeat = strings.TrimSpace(repeat)
	repeat = strings.ToLower(repeat)

	fmt.Println(repeat)
	weekRuleRegExp := `^w{1}\s+[1-7]{1}(,[1-7]){1,6}$`

	matched, err := regexp.MatchString(weekRuleRegExp, repeat)

	if err != nil {
		panic(err)
	}

	if matched {

		// дополнительно проверим уникальность дней, чтобы не было например такого: w 1,2,2,2,4
		re, _ := regexp.Compile(`\d{1}`)
		daysOfWeek := re.FindAllString(repeat, -1)
		if isUnique(daysOfWeek) {
			return true
		}
	}

	return false
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
