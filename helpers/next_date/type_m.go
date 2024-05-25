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

// возвращает новую дату выполнения задачи,
// если в repeat указаны только дни (аргумент days)
func calcMonthTypeWithDaysOnly(now, date time.Time, days string) (string, error) {
	daysList := strings.Split(days, ",")
	currentDate := date
	//fmt.Println(date)
	//date = date.AddDate(0, 1, 0)
	//fmt.Println(date)

	// Бесконечно бегаем по месяцам до тех пор, пока не поймаем нужный.
	// Начинаем с месяца из date
	for {

		// Получаем конкретные, упорядоченные по возрастанию, значения дней для каждого миесяца.
		// Ведь для разных месяцев дни, обозначающиеся
		// как -1 и -2, это разные дни.
		convertedDaysList := convertDays(date, daysList)

		fmt.Println(daysList)

		// бегаем по дням
		for i := 0; i < len(convertedDaysList); i++ {

			day := convertedDaysList[i]

			// если день из правила больше максимального дня рассматриваемом месяце - пропускаем
			if day > getLastDayOfMonth(date) {
				continue
			}

			// собираем потенциальную новую дату выполнения задачи
			newDate := time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, date.Location())

			//fmt.Println(newDate)

			// как только дата больше currentDate и now
			// возвращем её в нужном формате`
			if newDate.Sub(now) > 0 && newDate.Sub(currentDate) > 0 {
				return newDate.Format(DateFormat), nil
			}
		}

		// если мы пришли сюда, то прибавляем к date месяц и начинаем всё по-новой
		nextMonth := date.Month() + 1
		someDateOfNextMonth := time.Date(date.Year(), time.Month(nextMonth), 1, 0, 0, 0, 0, date.Location())
		lastDayOfNextMonth := getLastDayOfMonth(someDateOfNextMonth)

		// если последний день следующего месяца меньше последнего дня текущего месяца
		// и при этом в правиле есть символы -1 или -2
		// то дата следующего месяца = дата последнего дня следующего месяца
		if lastDayOfNextMonth < date.Day() && isSpecialDaySigns(daysList) {
			date = time.Date(date.Year(), time.Month(nextMonth), lastDayOfNextMonth, 0, 0, 0, 0, date.Location())
		} else {
			date = date.AddDate(0, 1, 0)
		}
	}
}

// проверяет на наличие специальных обозначений дней -1 или -2
func isSpecialDaySigns(dayList []string) bool {
	for i := 0; i < len(dayList); i++ {
		if dayList[i] == "-1" || dayList[i] == "-2" {
			return true
		}
	}

	return false
}

func convertMonths(months []string) []int {
	monthsInt := []int{}

	for i := 0; i < len(months); i++ {
		month, _ := strconv.Atoi(months[i])
		monthsInt = append(monthsInt, month)
	}

	sort.Slice(monthsInt, func(i, j int) bool {
		return monthsInt[i] < monthsInt[j]
	})

	return monthsInt
}

func calcMonthTypeWithDaysAndMonths(now, date time.Time, days, months string) (string, error) {

	// получим слайс месяцев в формате int
	monthsList := convertMonths(strings.Split(months, ","))
	currentDate := date

	// Бесконечно бегаем по годам до тех пор, пока не поймаем нужный.
	// Начинаем с года из date
	// на практике у этого цикла не должно быть больше 2 итераций
	for {
		// бегаем по месяцам
		for i := 0; i < len(monthsList); i++ {
			// someDate нужна только как аргумент в получении daysList
			someDate := time.Date(date.Year(), time.Month(monthsList[i]), 1, 0, 0, 0, 0, date.Location())
			daysList := convertDays(someDate, strings.Split(days, ","))

			// бегаем по дням
			for j := 0; j < len(daysList); j++ {
				day := daysList[j]

				// если день из правила больше максимального дня рассматриваемом месяце - пропускаем
				if day > getLastDayOfMonth(date) {
					continue
				}

				// собираем потенциальную новую дату выполнения задачи
				newDate := time.Date(date.Year(), time.Month(monthsList[i]), day, 0, 0, 0, 0, date.Location())

				// как только дата больше currentDate и now
				// возвращем её в нужном формате`
				if newDate.Sub(now) > 0 && newDate.Sub(currentDate) > 0 {
					return newDate.Format(DateFormat), nil
				}
			}
		}

		// если мы пришли сюда, то прибавляем к date год и начинаем всё по-новой
		date = date.AddDate(1, 0, 0)
	}
}

func calcMonthType(now time.Time, dateStr, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dateStr)

	if err != nil {
		return "", err
	}

	// разобьём строчку с правилом по пробелу
	repeatParts := strings.Fields(repeat)

	switch len(repeatParts) {
	case 2:
		return calcMonthTypeWithDaysOnly(now, date, repeatParts[1]) // расчёт без опциональной части с месяцами
	case 3:
		return calcMonthTypeWithDaysAndMonths(now, date, repeatParts[1], repeatParts[2]) // расчёт с опциональной частью
	default:
		return "", fmt.Errorf("переменная repeat недоспустимого формата: '" + repeat + "'")
	}
}