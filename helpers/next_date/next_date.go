package next_date

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const yearType = "y"
const dayType = "d"

// const weekType = "w"
const DateFormat = "20060102"

func Get(now time.Time, date string, repeat string) (string, error) {

	repeatTypeList := [2]string{yearType, dayType}
	var repeatType string
	for _, val := range repeatTypeList {
		if checkRepeatFormat(repeat, val) {
			repeatType = val
			fmt.Println(repeatType + " rule")
		}
	}

	switch repeatType {
	case yearType:
		return calcYearType(now, date)
	case dayType:
		return calcDayType(now, date, repeat)
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
