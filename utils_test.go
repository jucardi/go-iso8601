package iso8601

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMonthString(t *testing.T) {

}

func TestMonthMaps(t *testing.T) {
	maps := []map[time.Month]string{
		MonthsEng, MonthsEsp,
	}

	for _, m := range maps {
		assert.Len(t, m, 12)

		for i := 1; i <= 12; i++ {
			month := time.Month(i)
			_, ok := m[month]
			assert.True(t, ok, "Missing month ", month)

			// Testing long representation of months
			assert.Equal(t, m[month], GetMonthString(month, false, m), "Long string for month %s failed", month)

			// Testing short representation of months
			assert.Equal(t, m[month][0:3], GetMonthString(month, true, m), "Short string for month %s failed", month)
		}
	}
}

func TestTimeToString(t *testing.T) {
	mockTime := time.Time{}

	assert.Equal(t, "0001-01-01 00:00:00", TimeToString(mockTime, "yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "1-1-1 0:0:0", TimeToString(mockTime, "y-M-d H:m:s"))
	assert.Equal(t, "Jan 01, 0001", TimeToString(mockTime, "MMM dd, yyyy"))
	assert.Equal(t, "January 1, 01", TimeToString(mockTime, "MMMM d, yy"))
	assert.Equal(t, "Enero 1, 01 - 12:00:00AM", TimeToString(mockTime, "MMMM d, yy - hh:mm:sstt", MonthsEsp))
}
