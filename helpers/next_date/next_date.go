package next_date

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const yearType = "y"
const dayType = "d"
const weekType = "w"
const monthType = "m"
const typeCount = 4
const DateFormat = "20060102"

func Calc(now time.Time, date string, repeat string) (string, error) {

	repeatTypeList := [typeCount]string{yearType, dayType, weekType, monthType}
	var repeatType string
	for i := 0; i < typeCount; i++ {
		if checkRepeatFormat(repeat, repeatTypeList[i]) {
			repeatType = repeatTypeList[i]
			break
		}
	}

	switch repeatType {
	case yearType:
		return calcYearType(now, date)
	case dayType:
		return calcDayType(now, date, repeat)
	case weekType:
		return calcWeekType(now, date, repeat)
	case monthType:
		return calcMonthType(now, date, repeat)
	default:
		return "", fmt.Errorf("переменная repeat недоспустимого формата: '" + repeat + "'")
	}
}

func calcDayType(now time.Time, dateStr string, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dateStr)

	if err != nil {
		return "", err
	}

	re, _ := regexp.Compile(`\d+`)
	daysList := re.FindAllString(repeat, -1)
	days, err := strconv.Atoi(daysList[0])
	if err != nil {
		return "", err
	}

	newDate := date.AddDate(0, 0, days)

	for newDate.Sub(now) <= 0 {
		newDate = newDate.AddDate(0, 0, days)
	}

	fmt.Println(dateStr)

	return newDate.Format(DateFormat), nil
}

func calcYearType(now time.Time, dateStr string) (string, error) {

	date, err := time.Parse(DateFormat, dateStr)

	if err != nil {
		return "", err
	}

	newDate := date.AddDate(1, 0, 0)

	for newDate.Sub(now) <= 0 {
		newDate = newDate.AddDate(1, 0, 0)
	}

	return newDate.Format(DateFormat), nil
}
