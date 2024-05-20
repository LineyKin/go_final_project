package next_date

import (
	"fmt"
	"regexp"
	"strconv"
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
	case monthType:
		return checkMonthType(repeat)
	default:
		return false
	}
}

func checkMonthType(repeat string) bool {
	repeat = strings.TrimSpace(repeat)
	repeat = strings.ToLower(repeat)

	// разобьём строчку с правилом по пробелу
	repeatParts := strings.Fields(repeat)

	lenRepeatParts := len(repeatParts)

	// должно получиться 2 или 3 части. Если не так - проверка не пройдена
	if lenRepeatParts < 2 || lenRepeatParts > 3 {
		return false
	}

	// далее каждую часть анализируем отдельно

	// проверка первой обязательной части
	part1 := strings.TrimSpace(repeatParts[0])
	if part1 != monthType {
		fmt.Println("правило начинается с неверной буквы: " + part1)
		return false
	}

	// проверка второй обязательной части
	part2 := strings.TrimSpace(repeatParts[1])

	// проверим, что в строке нет однозначно неподходящих символов: только цифры, знак "," и знак "-"
	signsFilterRegExp := `^[\d,-]+$`
	matched, err := regexp.MatchString(signsFilterRegExp, part2)

	if err != nil {
		return false
	}

	if !matched {
		return false
	}

	// разбиваем строку по знаку "," чтобы получить слайс дней
	daysArr := strings.Split(part2, ",")

	// проверим уникальность: один день должен быть упомянут один раз
	if !isUnique(daysArr) {
		fmt.Println("дни не уникальны")
		return false
	}

	// день должен быть из множества [-2,0)U(0,31]
	// если хоть один день не попадает туда - проверка не пройдена]
	for i := 0; i < len(daysArr); i++ {
		day, _ := strconv.Atoi(daysArr[i])
		if day == 0 || day < -2 || day > 31 {
			fmt.Println("неверное значение дня: " + daysArr[i])
			return false
		}
	}

	// если опциональная часть отсустствует - проверка закончена успешно
	if lenRepeatParts == 2 {
		return true
	}

	// проверка третьей, опциональной, части
	part3 := repeatParts[2]

	// проверим, что в строке нет однозначно неподходящих символов: только цифры и знак ","
	signsFilterRegExp = `^[\d,]+$`
	matched, err = regexp.MatchString(signsFilterRegExp, part3)

	if err != nil {
		return false
	}

	if !matched {
		return false
	}

	monthArr := strings.Split(part3, ",")
	// проверим уникальность: один месяц должен быть упомянут один раз
	if !isUnique(monthArr) {
		fmt.Println("месяцы не уникальны")
		return false
	}

	for i := 0; i < len(monthArr); i++ {
		month, _ := strconv.Atoi(monthArr[i])
		if month < 1 || month > 12 {
			fmt.Println("неверное значение месяца: " + monthArr[i])
			return false
		}
	}

	// если мы дошли сюда, значит правило из трёх частей и все они прошли проверку
	return true
}

// проверяет отсустствие/присутствие повторяющихся элементов в слайсе
// если каждый элемент уникален - true, иначе false
func isUnique(arr []string) bool {
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
