package next_date

import (
	"fmt"
	//"strings"
	"time"
)

func Get(now time.Time, date string, repeat string) string {

	repeatTypeList := [4]string{"y", "d", "w", "m"}
	var repeatType string
	for _, val := range repeatTypeList {
		if checkRepeatFormat(repeat, val) {
			repeatType = val
			fmt.Println(repeatType + " rule")
		}
	}

	fmt.Println(now)
	return ""
}
