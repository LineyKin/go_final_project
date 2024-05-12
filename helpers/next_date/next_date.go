package next_date

import (
	"fmt"
	"time"
)

const yearType = "y"
const dayType = "d"
const dateFormat = "20060102"

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
	default:
		return "", fmt.Errorf("переменная repeat недоспустимого формата: '" + repeat + "'")
	}
}

func calcYearType(now time.Time, dateStr string) (string, error) {

	date, err := time.Parse(dateFormat, dateStr)

	if err != nil {
		return "", err
	}

	newDate := date.AddDate(1, 0, 0)

	for newDate.Sub(now) <= 0 {
		newDate = newDate.AddDate(1, 0, 0)
	}

	return newDate.Format(dateFormat), nil
}
