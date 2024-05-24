package next_date

import (
	"regexp"
	"strconv"
	"time"
)

// функция проверяет, подходит ли дата date под какой-нибудь
// день недели из списка weekDaysList
func checkWeekDay(date time.Time, weekDaysList []string) bool {
	weekDayToCheck := date.Weekday()
	for _, weekDayStr := range weekDaysList {
		weekDay, err := strconv.Atoi(weekDayStr)
		if err != nil {
			continue
		}

		// Go воскресенье понимает как 0, а не как 7
		if weekDay == 7 {
			weekDay = 0
		}

		if int(weekDayToCheck) == weekDay {
			return true
		}
	}

	return false
}

func calcWeekType(now time.Time, dateStr string, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dateStr)

	if err != nil {
		return "", err
	}

	re, _ := regexp.Compile(`\d+`)
	weekDaysList := re.FindAllString(repeat, -1)

	newDate := date.AddDate(0, 0, 1)

	for newDate.Sub(now) <= 0 || !checkWeekDay(newDate, weekDaysList) {
		newDate = newDate.AddDate(0, 0, 1)
	}

	return newDate.Format(DateFormat), nil
}
