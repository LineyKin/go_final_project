package next_date

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// выдаёт последнее число месяца по полной дате
func getLastDayOfMonth(date time.Time) int {
	firstOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return lastOfMonth.Day()
}

// придаёт конкретные значения чисел для дней, обозначающихся как "-1" и "-2"
// все дни переводятся в тип int
func convertDays(date time.Time, days []string) []int {
	daysInt := []int{}

	for i := 0; i < len(days); i++ {
		switch days[i] {
		case "-1":
			lastDay := getLastDayOfMonth(date)
			daysInt = append(daysInt, lastDay)
		case "-2":
			preLastDay := getLastDayOfMonth(date) - 1
			daysInt = append(daysInt, preLastDay)
		default:
			day, _ := strconv.Atoi(days[i])
			daysInt = append(daysInt, day)
		}
	}

	sort.Slice(daysInt, func(i, j int) bool {
		return daysInt[i] < daysInt[j]
	})

	return daysInt
}

func calcMonthTypeWithDaysOnly(now time.Time, date time.Time, days string) (string, error) {
	// прибавляем к дате месяц
	newDate := date.AddDate(0, 1, 0)

	daysList := strings.Split(days, ",")
	convertedDaysList := convertDays(newDate, daysList)
	fmt.Println(daysList)
	fmt.Println(convertedDaysList)

	return "", nil
}

func calcMonthType(now time.Time, dateStr string, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dateStr)

	if err != nil {
		return "", err
	}

	// разобьём строчку с правилом по пробелу
	repeatParts := strings.Fields(repeat)

	switch len(repeatParts) {
	case 2:
		return calcMonthTypeWithDaysOnly(now, date, repeatParts[1])
	case 3:
		return "both", nil
	default:
		return "", fmt.Errorf("переменная repeat недоспустимого формата: '" + repeat + "'")
	}
}
