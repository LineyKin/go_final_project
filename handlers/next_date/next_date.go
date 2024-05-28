package nextdate

import (
	nd "go_final_project/helpers/next_date"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// go test -run ^TestNextDate$ ./tests - OK
func GetNextDate(c *gin.Context) {
	now := c.Query("now")
	if now == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "нет параметра now"})
	}

	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "нет параметра date"})
	}

	repeat := c.Query("repeat")

	nowTime, _ := time.Parse(nd.DateFormat, now)
	nextDate, err := nd.Calc(nowTime, date, repeat)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, nextDate)
}
